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

type Project struct {
	ID           int64  `json:"id" db:"id,omitempty"`
	Name         string `json:"name" db:"name"`
	Info         string `json:"info" db:"info"`
	Ptype        string `json:"ptype" db:"ptype,omitempty"`
	OwnerID      int64  `json:"owner" db:"owner"`
	ExtraMeta    string `json:"extrameta" db:"extrameta,omitempty"`
	IsInitilized bool   `json:"isInitilized" db:"isInitilized,omitempty"`
}

type ProjectUser struct {
	ID          int64  `json:"id" db:"id,omitempty"`
	UserID      int64  `json:"userId" db:"userId"`
	ProjectID   int64  `json:"projectId" db:"projectId"`
	AccessToken string `json:"accessToken" db:"accessToken"`
	ExtraMeta   string `json:"extrameta" db:"extrameta,omitempty"`
}

type ProjectTypes struct {
	Name       string `json:"name"`
	Ptype      string `json:"ptype"`
	Info       string `json:"info"`
	Icon       string `json:"icon"`
	IsExternal bool   `json:"is_external"`
}
