package models

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

type Def struct {
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	Icon       string `json:"icon,omitempty"`
	Link       string `json:"link,omitempty"`
	IsExternal bool   `json:"is_external"`
}
