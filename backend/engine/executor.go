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
	Event         xhook.Event
	JsRuntime     *goja.Runtime
	ResultData    map[string]any
	PreventAction bool
}

// GOJA RUNTIME

func (e *Executor) executeJS(hook parsedHook) error {

	funcName := fmt.Sprintf("_handle_%d", hook.id)

	var entry func(ctx *goja.Object)
	eval := e.JsRuntime.Get(funcName)
	if eval == nil {
		return fmt.Errorf("%s function not found in script", funcName)
	}

	err := e.JsRuntime.ExportTo(eval, &entry)
	if err != nil {
		return err

	}

	obj := e.buildEventObject()

	entry(obj)

	return nil

}

func (e *Executor) buildEventObject() *goja.Object {
	obj := e.JsRuntime.NewObject()

	obj.Set("dataAsObject", func() any {

		return e.Event.Data
	})

	obj.Set("setPreventDefault", func() {
		e.PreventAction = true
	})

	obj.Set("event_name", e.Event.Name)
	obj.Set("event_id", e.Event.Id)
	obj.Set("project_id", e.Event.ProjectId)

	obj.Set("setResultData", func(data map[string]any) {
		e.ResultData = data
	})

	obj.Set("setResultDataField", func(field string, data any) {
		if e.ResultData == nil {
			e.ResultData = map[string]any{}
		}

		e.ResultData[field] = data
	})

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
	PreventAction bool           `json:"prevent_action"`
	ResultData    map[string]any `json:"result_data"`
}

func (e *Executor) executeWebhook(hook parsedHook) error {

	data := EventWebhookBody{
		Id:        e.Event.Id,
		Payload:   e.Event.Data,
		EventType: e.Event.Name,
		ProjectId: e.Event.ProjectId,
	}

	out, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	rawResp, err := http.Post(hook.target, "application/json", bytes.NewReader(out))
	if err != nil {
		return err
	}

	resp := &EventWebhookResponse{}

	err = json.NewDecoder(rawResp.Body).Decode(resp)
	if err != nil {
		return err
	}

	e.PreventAction = resp.PreventAction
	for key, val := range resp.ResultData {
		resp.ResultData[key] = val
	}

	return nil

}
