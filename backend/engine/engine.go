package engine

import (
	"errors"
	"fmt"
	"sync"

	"github.com/blue-monads/turnix/backend/xtypes"
	"github.com/blue-monads/turnix/backend/xtypes/models"
	"github.com/blue-monads/turnix/backend/xtypes/xproject"
	"github.com/k0kubun/pp"
)

type Options struct {
	App                xtypes.App
	Defs               map[string]*xproject.Defination
	ProjectInstallPath string
}

type Engine struct {
	app         xtypes.App
	globalJS    []byte
	projects    map[string]*xproject.Defination
	pLock       sync.RWMutex
	installPath string

	// control
	addPtypeChan    chan string
	removePtypeChan chan string
	updatePtypeChan chan string
}

func New(opts Options) *Engine {
	e := &Engine{
		projects:        opts.Defs,
		globalJS:        []byte(``),
		pLock:           sync.RWMutex{},
		app:             opts.App,
		addPtypeChan:    make(chan string),
		removePtypeChan: make(chan string),
		updatePtypeChan: make(chan string),
		installPath:     opts.ProjectInstallPath,
	}

	go e.eventLoop()

	return e
}

func (e *Engine) Start() error {
	e.load()

	return nil
}

func (e *Engine) eventLoop() {

	for {
		select {
		case ptype := <-e.addPtypeChan:
			pp.Println("@addPtypeChan", ptype)
		case ptype := <-e.updatePtypeChan:
			pp.Println("@updatePtypeChan", ptype)
		case ptype := <-e.removePtypeChan:
			pp.Println("@removePtypeChan", ptype)
		}
	}

}

func (e *Engine) InformPtypeAdded(ptype string) {
	e.addPtypeChan <- ptype
}

func (e *Engine) InformPtypeUpdated(ptype string) {
	e.updatePtypeChan <- ptype
}

func (e *Engine) InformPtypeRemoved(ptype string) {
	e.removePtypeChan <- ptype
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

func (e *Engine) InstallPath() string {
	return e.installPath
}
