package main

import (
	"encoding/json"
	"log/slog"
	"sync"

	"github.com/blue-monads/turnix/backend/modules/simplerat/wire"
	"github.com/mitchellh/mapstructure"
)

var (
	handlers = map[string]Handler{}
	hLock    = sync.Mutex{}
)

type Handler func(*WHContext)

func RegisterHandler(mtype string, handler Handler) {
	hLock.Lock()
	defer hLock.Unlock()
	handlers[mtype] = handler
}

// WHContext is the context for the worker handler
type WHContext struct {
	Packet     *wire.Packet
	Node       *AgentNode
	parsedData map[string]any
	pid        uint16
}

func (w *WHContext) Send(data []byte) {
	ch := w.Node.writeLoopCh

	ch <- &Response{
		Mid:        w.Packet.Header.Mid,
		Pid:        0,
		Mtype:      wire.PtypeResponseOk,
		DataBinary: data,
	}

}

func (w *WHContext) SendError(err string) {
	ch := w.Node.writeLoopCh

	ch <- &Response{
		Mid:        w.Packet.Header.Mid,
		Pid:        0,
		Mtype:      wire.PtypeResponseErr,
		DataBinary: []byte(err),
	}

}

func (w *WHContext) SendAsJSON(data any) error {
	b, err := json.Marshal(data)
	if err != nil {
		slog.Warn("SendAsJSON: Error marshaling data", slog.String("error", err.Error()))
		return err
	}

	w.Send(b)
	return nil
}

func (w *WHContext) SendPartial(data []byte) {
	ch := w.Node.writeLoopCh
	w.pid++
	pid := w.pid

	ch <- &Response{
		Mid:        w.Packet.Header.Mid,
		Pid:        pid,
		Mtype:      wire.PtypePartial,
		DataBinary: data,
	}

}

func (w *WHContext) SendPartialAsJSON(data any) error {
	b, err := json.Marshal(data)
	if err != nil {
		slog.Warn("SendPartialAsJSON: Error marshaling data", slog.String("error", err.Error()))
		return err
	}

	w.SendPartial(b)

	return nil
}

func (w *WHContext) init() {
	if w.parsedData != nil {
		return
	}

	err := json.Unmarshal(w.Packet.Data, &w.parsedData)
	if err != nil {
		slog.Warn("GetAsString: Error parsing data", slog.String("error", err.Error()))
	}

	w.Packet.Data = nil
}

func (w *WHContext) GetAsString(key string) string {
	w.init()

	val, ok := w.parsedData[key]
	if !ok {
		slog.Warn("GetAsString: Key not found", slog.Any("key", key))
		return ""
	}

	strVal, ok := val.(string)
	if ok {
		return strVal
	}

	return ""
}

func (w *WHContext) RootMap() map[string]any {
	w.init()
	return w.parsedData
}

func (w *WHContext) GetAsBool(key string) bool {

	w.init()

	val, ok := w.parsedData[key]
	if !ok {
		slog.Warn("GetAsBool: Key not found", slog.Any("key", key))
		return false
	}

	boolVal, ok := val.(bool)
	if ok {
		return boolVal
	}

	return false
}

func (w *WHContext) GetAsInt(key string) int {
	w.init()

	val, ok := w.parsedData[key]
	if !ok {
		slog.Warn("GetAsInt: Key not found", slog.Any("key", key))
		return 0
	}

	intVal, ok := val.(float64)
	if ok {
		return int(intVal)
	}

	return 0
}

func (w *WHContext) MapStructure(target any) error {
	if w.Packet.Data == nil {
		return json.Unmarshal(w.Packet.Data, &w.parsedData)
	}

	if w.parsedData != nil {
		return mapstructure.Decode(w.Packet.Data, target)
	}

	return nil

}
