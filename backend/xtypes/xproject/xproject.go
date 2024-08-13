package xproject

import (
	"io/fs"

	"github.com/bornjre/turnix/backend/xtypes"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type BuilderOption struct {
	App xtypes.App

	Logger zerolog.Logger
}

type Builder func(opt BuilderOption) (*Defination, error)

type Defination struct {
	Name                string
	Slug                string
	Info                string
	Icon                string
	Version             string
	NewFormSchemaFields []PTypeField
	Perminssions        []string
	EventTypes          []string
	GlobalJS            []byte
	AssetData           fs.FS

	OnAPIMount       func(RouterGroup *gin.RouterGroup)
	OnPageRequest    func(ctx *gin.Context)
	OnProjectRequest func(ctx *gin.Context)
	OnFileRequest    func(ctx *gin.Context)
	OnInit           func(pid int64) error
	IsInitilized     func(pid int64) (bool, error)
	OnDeInit         func(pid int64) error
}

type PTypeField struct {
	Name       string   `json:"name"`
	KeyName    string   `json:"key_name"`
	Ftype      string   `json:"ftype"`
	Disabled   bool     `json:"disabled"`
	Options    []string `json:"options"`
	IsRequired bool     `json:"is_required"`
}
