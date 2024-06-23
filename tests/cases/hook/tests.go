package hooktest

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"

	"github.com/bornjre/turnix/backend/distro"
	"github.com/bornjre/turnix/backend/xtypes/models"
	"github.com/bornjre/turnix/tests/must"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
)

func Run(app *distro.DistroApp, wg *sync.WaitGroup) {

	defer wg.Done()

	db := app.App.GetDatabase()

	users, err := db.ListUser()
	must.Handle(err)

	pp.Println(users)

	pid, err := db.AddProject(&models.Project{
		Name:         "test1",
		Ptype:        "test",
		OwnerID:      1,
		Info:         "this is test1",
		IsInitilized: true,
		IsPublic:     true,
	})

	pp.Println("@AddProject")

	must.Handle(err)

	_, err = db.AddProjectHook(1, pid, &models.ProjectHook{
		Name:     "sky_isfalling",
		Event:    EventSkyDropped,
		HookType: "script",
		HookCode: `
		
		const handle = (ctx) => {
			runtime.logInfo("I am calling from inside the hook js runtime :D", {})
			ctx.setResultDataField("invade_cuba_time", "11:30pm")
		}`,
		ProjectID: pid,
		Envs:      "{}",
	})

	pp.Println("@AddProjectHook")

	must.Handle(err)

	resp, err := http.Get("http://localhost:8777/z/project/test/xyz")
	must.Handle(err)

	if resp.StatusCode != 200 {
		pp.Println(resp.Status)
		panic("server error")
	}

	out, err := io.ReadAll(resp.Body)
	must.Handle(err)

	data := gin.H{}
	err = json.Unmarshal(out, &data)
	must.Handle(err)

}

/*

- test hook
- test dynamic asset serving
-


*/
