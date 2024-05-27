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
	IsEmailVerified bool       `json:"isEmailVerified" db:"email_verified"`
	ExtraMeta       string     `json:"extrameta" db:"extrameta,omitempty"`
	CreatedAt       *time.Time `json:"createdAt" db:"created_at,omitempty"`
	OwnerUserId     int64      `json:"owner_user_id" db:"owner_user_id,omitempty"`
	OwnerProjectId  int64      `json:"owner_project_id" db:"owner_project_id,omitempty"`
	MessageReadHead int64      `json:"msg_read_head" db:"msg_read_head,omitempty"`
	Disabled        bool       `json:"disabled" db:"disabled,omitempty"`
}

type UserMessage struct {
	ID            int64      `json:"id" db:"id,omitempty"`
	Name          string     `json:"title" db:"title"`
	Type          string     `json:"type" db:"type"`
	Contents      string     `json:"contents" db:"contents"`
	ToUser        int64      `json:"to_user" db:"to_user"`
	FromUser      int64      `json:"from_user" db:"from_user_id"`
	FromProject   int64      `json:"from_project" db:"from_project_id"`
	IsRead        bool       `json:"is_read" db:"is_read,omitempty"`
	CreatedAt     *time.Time `json:"created_at" db:"created_at,omitempty"`
	CallbackToken string     `json:"callback_token" db:"callback_token,omitempty"`
	Disabled      bool       `json:"disabled" db:"disabled,omitempty"`
}
