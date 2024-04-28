package models

type Project struct {
	ID           int64  `json:"id" db:"id,omitempty"`
	Name         string `json:"name" db:"name"`
	Info         string `json:"info" db:"info"`
	Ptype        string `json:"ptype" db:"ptype,omitempty"`
	OwnerID      int64  `json:"owner" db:"owned_by"`
	ExtraMeta    string `json:"extrameta" db:"extrameta,omitempty"`
	IsInitilized bool   `json:"is_initilized" db:"is_initilized,omitempty"`
	IsPublic     bool   `json:"is_public" db:"is_public,omitempty"`
}

type ProjectUser struct {
	ID        int64  `json:"id" db:"id,omitempty"`
	UserID    int64  `json:"user_id" db:"userId"`
	ProjectID int64  `json:"project_id" db:"projectId"`
	Scope     int64  `json:"scope" db:"scope,omitempty"`
	Token     string `json:"token" db:"token"`
	ExtraMeta string `json:"extrameta" db:"extrameta,omitempty"`
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
