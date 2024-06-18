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

type ProjectHook struct {
	ID          int64  `json:"id" db:"id,omitempty"`
	Name        string `json:"name" db:"name,omitempty"`
	Event       string `json:"event" db:"event"`
	OrderID     int64  `json:"order_id" db:"order_id"`
	RunasUserID int64  `json:"runas_user_id" db:"runas_user_id"`
	HookType    string `json:"hook_type" db:"hook_type"`
	HookCode    string `json:"hook_code" db:"hook_code"`
	Envs        string `json:"envs" db:"envs"`
	Target      string `json:"target" db:"target,omitempty"`
	ProjectID   int64  `json:"project_id" db:"project_id"`
	Extrameta   string `json:"extrameta" db:"extrameta,omitempty"`
	Priority    int64  `json:"priority" db:"priority,omitempty"`
}

type ProjectTypes struct {
	Name       string   `json:"name"`
	Ptype      string   `json:"ptype"`
	Slug       string   `json:"slug"`
	Info       string   `json:"info"`
	Icon       string   `json:"icon"`
	IsExternal bool     `json:"is_external"`
	Link       string   `json:"link,omitempty"`
	EventTypes []string `json:"event_types,omitempty"`
}
