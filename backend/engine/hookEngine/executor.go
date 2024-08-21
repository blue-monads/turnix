package hookengine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bornjre/turnix/backend/services/signer"
	"github.com/bornjre/turnix/backend/xtypes/xengine"
	"github.com/dop251/goja"
	"github.com/k0kubun/pp"
)

type ParsedHook struct {
	Id          int64
	Name        string
	EventType   string
	HookType    string
	RunasUserID int64
	Envs        map[string]string
	Target      string
	MustRun     bool
	MustNoError bool
}

type Executor struct {
	engine *HookEngine

	Event         xengine.EventContext
	JsRuntime     *goja.Runtime
	ResultData    map[string]any
	PreventAction bool
}

// GOJA RUNTIME

func (e *Executor) executeJS(hook ParsedHook) error {
	pp.Println("@executeJS")

	funcName := fmt.Sprintf("_handle_%d", hook.Id)

	var entry func(ctx *goja.Object)
	eval := e.JsRuntime.Get(funcName)
	if eval == nil {
		return fmt.Errorf("%s function not found in script", funcName)
	}

	err := e.JsRuntime.ExportTo(eval, &entry)
	if err != nil {
		return err
	}

	obj := e.buildEventObject(hook)

	pp.Println("@executeJS/entry")

	entry(obj)

	return nil

}

func (e *Executor) buildEventObject(hook ParsedHook) *goja.Object {
	obj := e.JsRuntime.NewObject()

	obj.Set("dataAsObject", func() any {

		return e.Event.Data
	})

	obj.Set("dataAsJsonString", func() any {

		out, err := json.Marshal(e.Event.Data)
		if err != nil {
			return nil
		}

		return string(out)
	})

	obj.Set("setPreventDefault", func() {
		e.PreventAction = true
	})

	obj.Set("event_type", e.Event.Event.Type)
	obj.Set("event_id", e.Event.EventId)
	obj.Set("project_id", e.Event.Event.Project)

	obj.Set("setResultData", func(data map[string]any) {
		e.ResultData = data
	})

	obj.Set("setResultDataField", func(field string, data any) {
		if e.ResultData == nil {
			e.ResultData = map[string]any{}
		}

		e.ResultData[field] = data
	})

	obj.Set("getEnv", func(field string) string {
		return hook.Envs[field]
	})

	obj.Set("getCtxUserToken", func() any {

		if hook.RunasUserID == -1 {
			return nil
		}

		userId := e.Event.Event.UserId
		if hook.RunasUserID != 0 {
			userId = hook.RunasUserID
		}

		token, err := e.engine.signer.SignAccess(&signer.AccessClaim{
			XID:    e.engine.snowflake.Generate().Int64(),
			Typeid: signer.TokenTypeAccess,
			UserId: userId,
			Extrameta: map[string]any{
				"engine_pid":      e.Event.Event.Project,
				"engine_event_id": e.Event.EventId,
			},
		})
		if err != nil {
			pp.Println("@couldnot/generateusertoken", err)
			return nil
		}

		return token

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

func (e *Executor) executeWebhook(hook ParsedHook) error {

	data := EventWebhookBody{
		Id:        e.Event.EventId,
		Payload:   e.Event.Data,
		EventType: e.Event.Event.Type,
		ProjectId: e.Event.Event.Project,
	}

	out, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	rawResp, err := http.Post(hook.Target, "application/json", bytes.NewReader(out))
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
