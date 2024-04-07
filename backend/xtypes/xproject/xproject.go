package xproject

import (
	"io/fs"

	"github.com/bornjre/trunis/backend/xtypes"
	"github.com/bornjre/trunis/backend/xtypes/xfiles"
	"github.com/bornjre/trunis/backend/xtypes/xsockd"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type BuilderOption struct {
	App xtypes.App

	Logger zerolog.Logger
}

type TypeBuilder func(opt BuilderOption) (ProjectType, error)

type TypeDefination struct {
	Name                string
	Slug                string
	Info                string
	Icon                string
	NewFormSchemaFields []PTypeField
	Perminssions        []string
	Builder             TypeBuilder
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

	HttpHandle(group *gin.RouterGroup)

	OnSockdMessage(msg *xsockd.Message) error
	OnSockdConn(opts *xsockd.ConnectOptions) error

	OnFileEvent(event *xfiles.Event) error
}
