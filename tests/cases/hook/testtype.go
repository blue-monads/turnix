package hooktest

import (
	"github.com/bornjre/turnix/backend/registry"
	"github.com/bornjre/turnix/backend/xtypes/services/xhook"
	"github.com/bornjre/turnix/backend/xtypes/xproject"
	"github.com/gin-gonic/gin"
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
		Builder:             New,
	})

}

type TestType struct {
}

func New(opt xproject.BuilderOption) (xproject.ProjectType, error) {

	t := &TestType{}

	opt.RouterGroup.GET("xyz", func(ctx *gin.Context) {
		engine := opt.App.GetHookEngine()

		result, err := engine.Emit(xhook.Event{
			Type:      EventSkyDropped,
			UserId:    1,
			ProjectId: 1,
			Data:      map[string]any{},
		})
		handle(err)

		itime := result.Data["invade_cuba_time"]
		pp.Println(result)

		if itime.(string) != "11:30pm" {
			panic("TEST FAILED")
		}

		ctx.Data(200, "application/json", []byte(`{ "hey": 123 }`))
	})

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
