package main

import (
	"log/slog"
	"sync"

	"github.com/mitchellh/mapstructure"
)

var (
	handlers = map[string]Handler{}
	hLock    = sync.Mutex{}
)

type Handler func(*WHContext) (any, error)

func RegisterHandler(mtype string, handler Handler) {
	hLock.Lock()
	defer hLock.Unlock()
	handlers[mtype] = handler
}

// WHContext is the context for the worker handler
type WHContext struct {
	Packet  *Packet
	Service *AgentService
}

func (w *WHContext) GetAsString(key string) string {
	maybeMap, ok := w.Packet.Data.Data.(map[string]any)
	if !ok {
		slog.Warn("GetAsString: Data is not a map", slog.Any("data", w.Packet.Data.Data))
		return ""
	}

	val, ok := maybeMap[key]
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
	m, ok := w.Packet.Data.Data.(map[string]any)
	if !ok {
		slog.Warn("RootMap: Data is not a map", slog.Any("data", w.Packet.Data.Data))
		return map[string]any{}
	}

	return m
}

func (w *WHContext) GetAsBool(key string) bool {
	maybeMap, ok := w.Packet.Data.Data.(map[string]any)
	if !ok {
		slog.Warn("GetAsBool: Data is not a map", slog.Any("data", w.Packet.Data.Data))
		return false
	}

	val, ok := maybeMap[key]
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
	maybeMap, ok := w.Packet.Data.Data.(map[string]any)
	if !ok {
		slog.Warn("GetAsInt: Data is not a map", slog.Any("data", w.Packet.Data.Data))
		return 0
	}

	val, ok := maybeMap[key]
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
	return mapstructure.Decode(w.Packet.Data.Data, target)
}
