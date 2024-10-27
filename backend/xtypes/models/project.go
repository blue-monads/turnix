package models

import "time"

type Project struct {
	ID           int64  `json:"id" db:"id,omitempty"`
	Name         string `json:"name" db:"name"`
	Info         string `json:"info" db:"info"`
	Ptype        string `json:"ptype" db:"ptype,omitempty"`
	OwnerID      int64  `json:"owned_by" db:"owned_by"`
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

type PluginImport struct {
	Name            string `json:"name" yaml:"name"`
	AppType         string `json:"apptype" yaml:"apptype"`
	ProjectTypeSlug string `json:"project_type_slug" yaml:"project_type_slug"`
	ServerCode      string `json:"server_code" yaml:"server_code"`
	ClientCode      string `json:"client_code" yaml:"client_code"`
}

type ProjectPlugin struct {
	ID         int64      `json:"id" db:"id,omitempty"`
	Name       string     `json:"name" db:"name"`
	Type       string     `json:"ptype" db:"ptype"`
	ProjectID  int64      `json:"project_id" db:"project_id"`
	ServerCode string     `json:"server_code" db:"server_code"`
	ClientCode string     `json:"client_code" db:"client_code"`
	CreatedBy  int64      `json:"created_by" db:"created_by"`
	UpdatedBy  int64      `json:"updated_by" db:"updated_by"`
	CreatedAt  *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at" db:"updated_at"`
}

type ProjectConfig struct {
	ID        int64  `json:"id" db:"id,omitempty"`
	Key       string `json:"key" db:"key"`
	Group     string `json:"group_name" db:"group_name"`
	Value     string `json:"value" db:"value"`
	ProjectID int64  `json:"project_id" db:"project_id"`
}

type TableColumn struct {
	Cid          int64  `json:"id" db:"id,omitempty"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	NotNull      bool   `json:"not_null"`
	DefaultValue string `json:"default_value"`
	PrimaryKey   bool   `json:"primary_key"`
}
