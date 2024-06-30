package hooktest

import (
	"github.com/bornjre/turnix/backend/registry"
	"github.com/bornjre/turnix/backend/xtypes/xbus"
	"github.com/bornjre/turnix/backend/xtypes/xproject"
	"github.com/bornjre/turnix/tests/must"
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
		engine := opt.App.GetEventBus()

		result, err := engine.Emit(xbus.EventNew{
			Type:    EventSkyDropped,
			UserId:  1,
			Project: 1,
			Data:    map[string]any{},
		})
		must.Handle(err)

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

func (t *TestType) Export(pid int64) (any, error) {
	return nil, nil
}

func (t *TestType) Restore(pid int64, data any) error {
	return nil
}
