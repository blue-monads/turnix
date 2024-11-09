package cmsengine

import "bytes"

type CMSEngine struct {
	actions map[string][]Handler
}

type ExecContext struct {
	e    *CMSEngine
	buf  bytes.Buffer
	next func() error
}

type ActionHandler func(e *ExecContext) error

type Handler func(e *CMSEngine) error
