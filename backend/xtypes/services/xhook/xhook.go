package xhook

import (
	"github.com/mitchellh/mapstructure"
)

type Engine interface {
	Init() error
	Invalidate(pid int64) error
	Emit(e Event) (*Result, error)
	Stop(force bool) error
}

type Event struct {
	Id        int64 // caller doesnot give it, its internal to runtime
	Type      string
	UserId    int64
	ProjectId int64
	Data      any
}

type Result struct {
	NoOfHooksRan  int16
	PreventAction bool
	Errors        map[string]string
	Data          map[string]any
}

// takes pointer to struct and maps fields bashed on json tags
func (r *Result) MapStruct(target any) error {

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "json",
		Result:  target,
	})
	if err != nil {
		return err
	}

	return decoder.Decode(r.Data)
}
