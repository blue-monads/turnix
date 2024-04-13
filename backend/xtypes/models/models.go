package models

import "time"

type User struct {
	ID              int64      `json:"id" db:"id,omitempty"`
	Name            string     `json:"name" db:"name"`
	Utype           string     `json:"utype" db:"utype,omitempty"`
	Email           string     `json:"email" db:"email,omitempty"`
	Phone           string     `json:"phone" db:"phone,omitempty"`
	Bio             string     `json:"bio" db:"bio"`
	Password        string     `json:"password" db:"password,omitempty"`
	IsEmailVerified bool       `json:"isEmailVerified" db:"isEmailVerified"`
	ExtraMeta       string     `json:"extrameta" db:"extrameta,omitempty"`
	CreatedAt       *time.Time `json:"createdAt" db:"createdAt,omitempty"`
	OwnedBy         int64      `json:"ownedBy" db:"ownedBy,omitempty"`
	AccessToken     string     `json:"accessToken" db:"accessToken,omitempty"`
	Disabled        string     `json:"disabled" db:"disabled,omitempty"`
}
