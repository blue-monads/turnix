package hookengine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bornjre/turnix/backend/xtypes/services/xhook"
	"github.com/dop251/goja"
)

type Executor struct {
	evt           xhook.Event
	runner        *hookRunner
	jrt           *goja.Runtime
	preventAction bool
}

// GOJA RUNTIME

func (e *Executor) executeJS(hook parsedHook) error {

	funcName := fmt.Sprintf("_handle_%d", hook.id)

	var entry func(ctx *goja.Object)
	eval := e.jrt.Get(funcName)
	if eval == nil {
		return fmt.Errorf("%s function not found in script", funcName)
	}

	err := e.jrt.ExportTo(eval, &entry)
	if err != nil {
		return err

	}

	obj := buildEventObject(e.evt, e.jrt)

	entry(obj)

	return nil

}

func buildEventObject(evt xhook.Event, r *goja.Runtime) *goja.Object {
	obj := r.NewObject()

	obj.Set("dataAsObject", func() any {

		return evt.Data
	})

	obj.Set("name", evt.Name)

	return obj

}

// WEBHOOK

type EventWebhookBody struct {
	Id        int64  `json:"xid"`
	Payload   any    `json:"payload"`
	EventType string `json:"event_type"`
	ProjectId int64  `json:"project_id"`
}

type EventWebhookResponse struct {
	PreventAction bool `json:"prevent_action"`
	NewPayload    any  `json:"new_payload"`
}

func (e *Executor) executeWebhook(hook parsedHook) error {

	data := EventWebhookBody{
		Id:        1,
		Payload:   e.evt.Data,
		EventType: e.evt.Name,
		ProjectId: e.evt.ProjectId,
	}

	out, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	_, err = http.Post(hook.target, "application/json", bytes.NewReader(out))
	if err != nil {
		return err
	}

	return nil

}
