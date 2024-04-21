package models

import (
	"time"
)

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
	Disabled        bool       `json:"disabled" db:"disabled,omitempty"`
}

type UserMessage struct {
	ID            int64      `json:"id" db:"id,omitempty"`
	Name          string     `json:"title" db:"title"`
	Type          string     `json:"type" db:"type"`
	Contents      string     `json:"contents" db:"contents"`
	ToUser        int64      `json:"to_user" db:"toUser"`
	FromUser      int64      `json:"from_user" db:"fromUser"`
	FromProject   int64      `json:"from_project" db:"fromProject"`
	IsRead        bool       `json:"is_read" db:"isRead,omitempty"`
	CreatedAt     *time.Time `json:"createdAt" db:"createdAt,omitempty"`
	CallbackToken string     `json:"callbackToken" db:"callbackToken,omitempty"`
	Disabled      bool       `json:"disabled" db:"disabled,omitempty"`
}
