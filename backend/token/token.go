package token

import (
	"crypto/sha256"
	"encoding/json"
	"errors"

	"github.com/hako/branca"
	"golang.org/x/crypto/pbkdf2"
)

const (
	TokenTypeAccess      uint8 = 1
	TokenTypeEmailInvite uint8 = 2
)

type AccessClaim struct {
	XID    int64 `json:"x,omitempty"`
	Typeid uint8 `json:"t,omitempty"`
	UserId int64 `json:"u,omitempty"`
}

type InviteClaim struct {
	XID            int64  `json:"x,omitempty"`
	Typeid         uint8  `json:"t,omitempty"`
	InviteFromUser int64  `json:"u,omitempty"`
	InviteToUser   int64  `json:"z,omitempty"`
	InviteEmail    string `json:"e,omitempty"`
	ProjectId      int64  `json:"p,omitempty"`
}

// fixme => add expiry

var ErrInvalidToken = errors.New("INVALID TOKEN")

type TokenService struct {
	signer *branca.Branca
}

func New(key []byte) *TokenService {
	masterKey := pbkdf2.Key(key, []byte("SALTY_SALMON"), 2048, 32, sha256.New)

	return &TokenService{
		signer: branca.NewBranca(string(masterKey)),
	}
}

func (t *TokenService) parse(token string, dest any) error {
	str, err := t.signer.DecodeToString(token)
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(str), dest)
}

func (t *TokenService) sign(o any) (string, error) {
	out, err := json.Marshal(o)
	if err != nil {
		return "", nil
	}

	return t.signer.EncodeToString(string(out))
}

func (ts *TokenService) ParseAccess(tstr string) (*AccessClaim, error) {

	claim := &AccessClaim{}

	err := ts.parse(tstr, claim)
	if err != nil {
		return nil, err
	}

	if claim.Typeid != TokenTypeAccess {
		return nil, ErrInvalidToken
	}

	return claim, nil
}

func (ts *TokenService) SignAccess(claim *AccessClaim) (string, error) {

	claim.Typeid = TokenTypeAccess

	return ts.sign(claim)
}

func (ts *TokenService) ParseInvite(tstr string) (*InviteClaim, error) {

	claim := &InviteClaim{}

	err := ts.parse(tstr, claim)
	if err != nil {
		return nil, err
	}

	if claim.Typeid != TokenTypeEmailInvite {
		return nil, ErrInvalidToken
	}

	return claim, nil
}

func (ts *TokenService) SignInvite(claim *InviteClaim) (string, error) {

	claim.Typeid = TokenTypeEmailInvite

	return ts.sign(claim)
}
