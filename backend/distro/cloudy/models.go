package cloudy

import "time"

type User struct {
	ID           int64      `db:"id,omitempty" json:"id"`
	Fullname     string     `db:"fullname" json:"fullname"`
	Email        string     `db:"email" json:"email"`
	Password     string     `db:"password" json:"-"`
	TenantKey    string     `db:"tenant_key" json:"tenant_key"`
	UType        string     `db:"utype" json:"utype"`
	PricingTier  string     `db:"pricing_tier" json:"pricing_tier"`
	IsVerified   bool       `db:"is_verified" json:"is_verified"`
	IsLazyLoaded bool       `db:"is_lazy_loaded" json:"is_lazy_loaded"`
	IsDisabled   bool       `db:"is_disabled" json:"is_disabled"`
	CreatedAt    *time.Time `db:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt    *time.Time `db:"updated_at,omitempty" json:"updated_at,omitempty"`
}
