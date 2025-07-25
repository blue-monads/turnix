package engine

import (
	"errors"
	"fmt"
	"sync"

	"github.com/blue-monads/turnix/backend/xtypes"
	"github.com/blue-monads/turnix/backend/xtypes/models"
	"github.com/blue-monads/turnix/backend/xtypes/xproject"
	"github.com/k0kubun/pp"
	"github.com/rs/zerolog"
)

type Options struct {
	App                xtypes.App
	Defs               map[string]*xproject.Defination
	BasePath           string
	ProjectInstallPath string
	Logger             zerolog.Logger
}

type Engine struct {
	opts     Options
	globalJS []byte
	projects map[string]*LoadedDef
	pLock    sync.RWMutex
	// control
	addPtypeChan    chan string
	removePtypeChan chan string
	updatePtypeChan chan string
}

func New(opts Options) *Engine {
	defs := make(map[string]*LoadedDef)

	for k, v := range opts.Defs {
		ldef := &LoadedDef{
			def:     v,
			defType: defTypeNative,
		}

		defs[k] = ldef
	}

	e := &Engine{
		projects:        defs,
		globalJS:        []byte(``),
		pLock:           sync.RWMutex{},
		addPtypeChan:    make(chan string),
		removePtypeChan: make(chan string),
		updatePtypeChan: make(chan string),
		opts:            opts,
	}

	go e.eventLoop()

	return e
}

func (e *Engine) Start() error {
	e.load()

	return nil
}

func (e *Engine) eventLoop() {

	loadApp := func(ptype string) {
		e.pLock.RLock()
		ldef := e.projects[ptype]
		e.pLock.RUnlock()

		if ldef == nil {
			return
		}

		if ldef.defType == defTypeNative {
			// cannot reload native type
			return
		}

		if ldef.def.OnClose != nil {
			ldef.def.OnClose()
		}

		if ldef.defType == defTypeLuaZip {
			e.LoadPtypeWithZip(ldef.file)
		}

		if ldef.defType == defTypeLuaFolder {
			e.LoadPtypeWithFolder(ldef.file)
		}

	}

	unloadApp := func(ptype string) {
		e.pLock.Lock()
		ldef := e.projects[ptype]
		delete(e.projects, ptype)
		e.pLock.Unlock()

		if ldef == nil {
			return
		}

		if ldef.defType == defTypeNative {
			// cannot unload native type
			return
		}

		if ldef.def.OnClose != nil {
			ldef.def.OnClose()
		}

	}

	for {
		select {
		case ptype := <-e.addPtypeChan:
			pp.Println("@addPtypeChan", ptype)
			loadApp(ptype)
		case ptype := <-e.updatePtypeChan:
			pp.Println("@updatePtypeChan", ptype)
			loadApp(ptype)
		case ptype := <-e.removePtypeChan:
			pp.Println("@removePtypeChan", ptype)
			unloadApp(ptype)
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
			Name:               pdef.def.Name,
			Ptype:              pdef.def.Slug,
			Icon:               pdef.def.Icon,
			Info:               pdef.def.Info,
			IsExternal:         pdef.def.LinkPattern != "",
			Slug:               pdef.def.Slug,
			ProjectLinkPattern: pdef.def.LinkPattern,
			BaseLink:           fmt.Sprintf("/z/pages/portal/projects/%s", pdef.def.Slug),
		})

	}

	return pdefs, nil
}

func (e *Engine) GetProjectType(ptype string) (*models.ProjectTypes, error) {

	for _, pdef := range e.projects {

		if pdef.def.Slug == ptype {
			return &models.ProjectTypes{
				Name:       pdef.def.Name,
				Ptype:      pdef.def.Slug,
				Slug:       pdef.def.Slug,
				Info:       pdef.def.Info,
				Icon:       pdef.def.Icon,
				IsExternal: pdef.def.AssetData != nil,
			}, nil
		}

	}

	return nil, errors.New("not found")
}

func (e *Engine) GetProjectTypeReload(ptype string) error {
	e.InformPtypeUpdated(ptype)
	return nil
}

func (e *Engine) GetProjectTypeForm(ptype string) ([]xproject.PTypeField, error) {

	for _, pdef := range e.projects {
		if pdef.def.Slug == ptype {
			return pdef.def.NewFormSchemaFields, nil
		}
	}

	return nil, errors.New("not found")

}

func (e *Engine) InstallPath() string {
	return e.opts.ProjectInstallPath
}
