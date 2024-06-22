package main

import (
	"github.com/bornjre/turnix/backend/distro"
	"github.com/bornjre/turnix/backend/xtypes/models"
)

func runTest(app *distro.DistroApp) {

	db := app.App.GetDatabase()

	pid, err := db.AddProject(&models.Project{
		Name:         "test1",
		Ptype:        "test",
		OwnerID:      1,
		Info:         "this is test1",
		IsInitilized: true,
		IsPublic:     true,
	})

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

	handle(err)

}
