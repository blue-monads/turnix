package engine

import (
	"errors"
	"fmt"
	"sync"

	"github.com/blue-monads/turnix/backend/xtypes"
	"github.com/blue-monads/turnix/backend/xtypes/models"
	"github.com/blue-monads/turnix/backend/xtypes/xproject"
)

type Options struct {
	App  xtypes.App
	Defs map[string]*xproject.Defination
}

type Engine struct {
	app      xtypes.App
	globalJS []byte
	projects map[string]*xproject.Defination
	pLock    sync.RWMutex
}

func New(opts Options) *Engine {
	return &Engine{
		projects: opts.Defs,
		globalJS: []byte(``),
		pLock:    sync.RWMutex{},
		app:      opts.App,
	}
}

func (e *Engine) ListProjectTypes() ([]models.ProjectTypes, error) {
	pdefs := make([]models.ProjectTypes, 0)

	for _, pdef := range e.projects {
		pdefs = append(pdefs, models.ProjectTypes{
			Name:               pdef.Name,
			Ptype:              pdef.Slug,
			Icon:               pdef.Icon,
			Info:               pdef.Info,
			IsExternal:         pdef.LinkPattern != "",
			Slug:               pdef.Slug,
			ProjectLinkPattern: pdef.LinkPattern,
			BaseLink:           fmt.Sprintf("/z/pages/portal/projects/%s", pdef.Slug),
		})

	}

	return pdefs, nil
}

func (e *Engine) GetProjectType(ptype string) (*models.ProjectTypes, error) {

	for _, pdef := range e.projects {

		if pdef.Slug == ptype {
			return &models.ProjectTypes{
				Name:       pdef.Name,
				Ptype:      pdef.Slug,
				Slug:       pdef.Slug,
				Info:       pdef.Info,
				Icon:       pdef.Icon,
				IsExternal: pdef.AssetData != nil,
			}, nil
		}

	}

	return nil, errors.New("not found")
}

func (e *Engine) GetProjectTypeForm(ptype string) ([]xproject.PTypeField, error) {

	for _, pdef := range e.projects {
		if pdef.Slug == ptype {
			return pdef.NewFormSchemaFields, nil
		}
	}

	return nil, errors.New("not found")

}
