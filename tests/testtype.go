package main

import (
	"github.com/bornjre/turnix/backend/registry"
	"github.com/bornjre/turnix/backend/xtypes"
	"github.com/bornjre/turnix/backend/xtypes/services/xhook"
	"github.com/bornjre/turnix/backend/xtypes/xproject"
	"github.com/k0kubun/pp"
)

const (
	EventSkyDropped = "sky_dropped"
)

func init() {
	registry.Register(&xproject.Defination{
		Name:                "test",
		Slug:                "test",
		Info:                "this is test type project",
		NewFormSchemaFields: []xproject.PTypeField{},
		Perminssions:        []string{},
		EventTypes:          []string{EventSkyDropped},
		Builder:             nil,
	})

}

type TestType struct {
}

func New(opt xproject.BuilderOption) (xproject.ProjectType, error) {

	t := &TestType{}

	opt.RouterGroup.GET("testxyz", opt.App.AuthMiddleware(func(ctx xtypes.ContextPlus) (any, error) {

		engine := opt.App.GetHookEngine()

		result, err := engine.Emit(xhook.Event{
			Name:      EventSkyDropped,
			UserId:    ctx.Claim.UserId,
			ProjectId: ctx.ParamInt64("pid"),
			Data:      map[string]any{},
		})

		if err != nil {
			return nil, err
		}

		pp.Println("@result", result)

		return nil, nil
	}))

	return t, nil
}

func (t *TestType) Init(pid int64) error {

	return nil
}

func (t *TestType) IsInitilized(pid int64) (bool, error) {

	return true, nil
}

func (t *TestType) DeInit(pid int64) error {
	return nil
}

func (t *TestType) OnFileEvent(event *xproject.FileEvent) error {
	return nil
}

func (t *TestType) OnUserEvent(event *xproject.UserEvent) error {
	return nil
}
