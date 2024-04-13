package xfilestore

import "github.com/gin-gonic/gin"

type EventType = uint8

const (
	EventTypeList EventType = iota
	EventTypeAddFile
	EventTypeAddFolder
)

type Event struct {
	UserId      int64
	ProjectId   int64
	Type        EventType
	Path        string
	Action      string
	HttpContext *gin.Context
	Consumed    bool
}
