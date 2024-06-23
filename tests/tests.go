package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/bornjre/turnix/backend/distro"
	"github.com/bornjre/turnix/backend/xtypes/models"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
)

func runTest(app *distro.DistroApp) {

	pp.Println("@runTest/sleeping/before")

	time.Sleep(time.Second * 3)

	pp.Println("@runTest/sleeping/end")

	db := app.App.GetDatabase()

	users, err := db.ListUser()
	handle(err)

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

	handle(err)

	_, err = db.AddProjectHook(1, pid, &models.ProjectHook{
		Name:     "sky_isfalling",
		Event:    EventSkyDropped,
		HookType: "script",
		HookCode: `
		
		const handle = (ctx) => {
			runtime.logInfo("I am calling from inside the hook js runtime :D", {})
		}`,
		ProjectID: pid,
		Envs:      "{}",
	})

	pp.Println("@AddProjectHook")

	handle(err)

	resp, err := http.Get("http://localhost:8777/z/project/test/xyz")
	handle(err)

	if resp.StatusCode != 200 {
		pp.Println(resp.Status)
		panic("server error")
	}

	out, err := io.ReadAll(resp.Body)
	handle(err)

	data := gin.H{}
	err = json.Unmarshal(out, &data)
	handle(err)

}

/*

- test hook
- test dynamic asset serving
-


*/
