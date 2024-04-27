package xproject

import (
	"io/fs"

	"github.com/bornjre/trunis/backend/xtypes"
	"github.com/bornjre/trunis/backend/xtypes/services/xsockd"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type BuilderOption struct {
	App xtypes.App

	Logger zerolog.Logger

	RouterGroup *gin.RouterGroup
}

type Builder func(opt BuilderOption) (ProjectType, error)

type Defination struct {
	Name                string
	Slug                string
	Info                string
	Icon                string
	NewFormSchemaFields []PTypeField
	Perminssions        []string
	Builder             Builder
	GlobalJS            []byte
	AssetData           fs.FS
	InterceptFileEvent  bool
}

type PTypeField struct {
	Name       string   `json:"name"`
	KeyName    string   `json:"key_name"`
	Ftype      string   `json:"ftype"`
	Disabled   bool     `json:"disabled"`
	Options    []string `json:"options"`
	IsRequired bool     `json:"is_required"`
}

type ProjectType interface {
	Init(pid int64) error

	IsInitilized(pid int64) (bool, error)

	DeInit(pid int64) error

	OnSockdMessage(msg *xsockd.Message) error
	OnSockdConn(opts *xsockd.ConnectOptions) error

	OnFileEvent(event *FileEvent) error

	OnUserEvent(event *UserEvent) error
}
