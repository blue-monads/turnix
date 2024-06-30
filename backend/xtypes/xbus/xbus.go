package xbus

import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type EventNew struct {

	// optional userid associated to the session
	UserId int64

	// optional project id if it is generated from project not system event
	Project int64

	// event type
	Type string

	// optional path of the entitity this event is generated from
	Path string

	// actual event payloa
	Data any

	// allow hook to be executed
	AllowHook bool

	// optional http context
	HttpContext *gin.Context
}

type EventResult struct {
	PreventAction bool

	Errors map[string]string
	Data   map[string]any
}

type EventContext struct {
	EventId       int64
	Event         *EventNew
	PreventAction bool
	Data          map[string]any
}

// map struct data to struct, target to pointer to struct
func (r *EventContext) MapStruct(target any) error {

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "json",
		Result:  target,
	})
	if err != nil {
		return err
	}

	return decoder.Decode(r.Data)
}

type EventHandler func(ctx EventContext) error

type EventBus interface {
	Init() error
	Invalidate(pid int64) error
	Emit(e EventNew) (*EventResult, error)

	OnEvent(name string, handler EventHandler)

	Stop(force bool) error
}
