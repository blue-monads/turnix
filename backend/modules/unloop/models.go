package unloop

import "time"

type Template struct {
	ID            int64  `json:"id" db:"id,omitempty"`
	Name          string `json:"name" db:"name"`
	Ttype         string `json:"ttype" db:"ttype,omitempty"`
	CanReadResult string `json:"canReadResult" db:"canReadResult"`
	Consts        string `json:"consts" db:"consts"`
	ContentScript string `json:"contentScript" db:"contentScript"`
	ProjectID     int64  `json:"projectId" db:"projectId"`
	ExtraMeta     string `json:"extrameta" db:"extrameta,omitempty"`
}

type Message struct {
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
