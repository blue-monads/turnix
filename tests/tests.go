package main

import (
	"time"

	"github.com/bornjre/turnix/backend/distro"
	"github.com/bornjre/turnix/backend/xtypes/models"
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
			runtime.logInfo("This is test", {})
		}`,
		ProjectID: pid,
		Envs:      "{}",
	})

	pp.Println("@AddProjectHook")

	handle(err)

}
