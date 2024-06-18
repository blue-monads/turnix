package xproject

import (
	"github.com/gin-gonic/gin"
)

type EventType = uint16

const (
	EventTypeList EventType = iota
	EventTypeAddFile
	EventTypeAddFolder
	EventTypeUserLogin
	EventTypeHook
)

type FileEvent struct {
	UserId      int64
	ProjectId   int64
	Type        EventType
	Path        string
	Action      string
	HttpContext *gin.Context
	Consumed    bool
}

type UserEvent struct {
	UserId      int64
	ProjectId   int64
	Type        EventType
	Path        string
	Action      string
	HttpContext *gin.Context
	Consumed    bool
}
