package database

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
}

type Project struct {
	ID        int64  `djson:"id" db:"id,omitempty"`
	Ptype     string `json:"ptype" db:"ptype,omitempty"`
	Name      string `json:"name" db:"name"`
	OwnerID   int64  `json:"owner" db:"owner"`
	ExtraMeta string `json:"extrameta" db:"extrameta,omitempty"`
}

type ProjectUser struct {
	ID          int64  `json:"id" db:"id,omitempty"`
	UserID      int64  `json:"userId" db:"userId"`
	ProjectID   int64  `json:"projectId" db:"projectId"`
	AccessToken string `json:"accessToken" db:"accessToken"`
	ExtraMeta   string `json:"extrameta" db:"extrameta,omitempty"`
}

type PTLactionTemplate struct {
	ID            int64  `json:"id" db:"id,omitempty"`
	Name          string `json:"name" db:"name"`
	Ttype         string `json:"ttype" db:"ttype,omitempty"`
	CanReadResult string `json:"canReadResult" db:"canReadResult"`
	Consts        string `json:"consts" db:"consts"`
	ContentScript string `json:"contentScript" db:"contentScript"`
	ProjectID     int64  `json:"projectId" db:"projectId"`
	ExtraMeta     string `json:"extrameta" db:"extrameta,omitempty"`
}

type PTLactionQueueMessage struct {
	ID         int64      `json:"id" db:"id,omitempty"`
	Submitter  int64      `json:"submitter" db:"submitter"`
	ProjectID  int64      `json:"projectId" db:"projectId"`
	TemplateID int64      `json:"templateId" db:"templateId"`
	Contents   string     `json:"contents" db:"contents"`
	Result     string     `json:"result" db:"result"`
	Status     string     `json:"status" db:"status"`
	ExtraMeta  string     `json:"extrameta" db:"extrameta"`
	CreatedAt  *time.Time `json:"createdAt" db:"createdAt,omitempty"`
}
