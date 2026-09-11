package cloudy

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/sqlite"
)

type Store struct {
	sess db.Session
	sql  *sql.DB
}

func newStore(sqlDB *sql.DB) (*Store, error) {
	sess, err := sqlite.New(sqlDB)
	if err != nil {
		return nil, err
	}
	return &Store{sess: sess, sql: sqlDB}, nil
}

func (s *Store) users() db.Collection {
	return s.sess.Collection("CloudyUsers")
}

func (s *Store) listUsers() ([]*User, error) {
	var users []*User
	err := s.users().Find().OrderBy("id").All(&users)
	if err != nil {
		return nil, err
	}
	for _, u := range users {
		normalizeUser(u)
	}
	return users, nil
}

func (s *Store) getUserByID(id int64) (*User, error) {
	u := &User{}
	err := s.users().Find(db.Cond{"id": id}).One(u)
	if errors.Is(err, db.ErrNoMoreRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	normalizeUser(u)
	return u, nil
}

func (s *Store) getUserByEmail(email string) (*User, error) {
	u := &User{}
	err := s.users().Find(db.Cond{"email": email}).One(u)
	if errors.Is(err, db.ErrNoMoreRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	normalizeUser(u)
	return u, nil
}

func (s *Store) getUserByTenant(tenantKey string) (*User, error) {
	u := &User{}
	err := s.users().Find(db.Cond{"tenant_key": tenantKey}).One(u)
	if errors.Is(err, db.ErrNoMoreRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	normalizeUser(u)
	return u, nil
}

func (s *Store) insertUser(fullname, email, passwordHash, tenantKey, pricingTier, utype string, verified bool) (*User, error) {
	if utype == "" {
		utype = UTypeNormal
	}
	if !validUType(utype) {
		return nil, fmt.Errorf("utype must be admin or normal")
	}
	if pricingTier == "" {
		pricingTier = "free"
	}

	now := time.Now().UTC()
	u := &User{
		Fullname:     fullname,
		Email:        email,
		Password:     passwordHash,
		TenantKey:    tenantKey,
		UType:        utype,
		PricingTier:  pricingTier,
		IsVerified:   verified,
		IsLazyLoaded: true,
		IsDisabled:   false,
		CreatedAt:    &now,
		UpdatedAt:    &now,
	}

	res, err := s.users().Insert(u)
	if err != nil {
		return nil, err
	}
	id, ok := res.ID().(int64)
	if !ok {
		// sqlite sometimes returns int
		switch v := res.ID().(type) {
		case int:
			id = int64(v)
		default:
			return nil, fmt.Errorf("unexpected insert id type %T", res.ID())
		}
	}
	return s.getUserByID(id)
}

func (s *Store) markUserVerified(id int64) error {
	now := time.Now().UTC()
	return s.users().Find(db.Cond{"id": id}).Update(map[string]any{
		"is_verified": true,
		"updated_at":  now,
	})
}

func (s *Store) updateUserPassword(id int64, passwordHash string) error {
	now := time.Now().UTC()
	return s.users().Find(db.Cond{"id": id}).Update(map[string]any{
		"password":   passwordHash,
		"updated_at": now,
	})
}

func (s *Store) setUserDisabled(id int64, disabled bool) error {
	now := time.Now().UTC()
	return s.users().Find(db.Cond{"id": id}).Update(map[string]any{
		"is_disabled": disabled,
		"updated_at":  now,
	})
}

func (s *Store) setUserLazyLoaded(id int64, lazy bool) error {
	now := time.Now().UTC()
	return s.users().Find(db.Cond{"id": id}).Update(map[string]any{
		"is_lazy_loaded": lazy,
		"updated_at":     now,
	})
}

func (s *Store) tenantExists(tenantKey string) bool {
	u, err := s.getUserByTenant(tenantKey)
	return err == nil && u != nil
}

func normalizeUser(u *User) {
	if u == nil {
		return
	}
	if u.UType == "" {
		u.UType = UTypeNormal
	}
}
