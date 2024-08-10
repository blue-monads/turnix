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

	RouterGroup *gin.RouterGroup
}

type Builder func(opt BuilderOption) (ProjectType, error)

type Defination struct {
	Name                string
	Slug                string
	Info                string
	Icon                string
	Version             string
	NewFormSchemaFields []PTypeField
	Perminssions        []string
	EventTypes          []string
	Builder             Builder
	GlobalJS            []byte
	AssetData           fs.FS

	OnPageRequest func(ctx *gin.Context) (bool, error)
	OnFileRequest func(ctx *gin.Context) (bool, error)
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
}
