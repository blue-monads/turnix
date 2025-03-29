package engine

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/blue-monads/turnix/backend/engine/luaz"
	"github.com/blue-monads/turnix/backend/xtypes/models"
	"github.com/blue-monads/turnix/backend/xtypes/xproject"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
)

const (
	defTypeNative    uint8 = 0
	defTypeLuaZip    uint8 = 1
	defTypeLuaFolder uint8 = 2
)

type LoadedDef struct {
	def     *xproject.Defination
	defType uint8
	file    string
}

func (e *Engine) load() {

	files, err := os.ReadDir(e.opts.ProjectInstallPath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		pp.Println(file.Name())

		if file.IsDir() {
			e.LoadPtypeWithFolder(path.Join(e.opts.ProjectInstallPath, file.Name()))
		}

		if strings.HasSuffix(file.Name(), ".zip") {
			e.LoadPtypeWithZip(path.Join(e.opts.ProjectInstallPath, file.Name()))
		}

	}

	pp.Println("loaded", len(e.projects))
}

func (e *Engine) LoadPtypeWithFolder(filePath string) error {

	pp.Println("LoadPtypeWithFolder/1")

	// Read the manifest file
	manifestFile := path.Join(filePath, manifestFileName)
	fbytes, err := os.ReadFile(manifestFile)
	if err != nil {
		return err
	}

	pp.Println("LoadPtypeWithFolder/2")

	// Parse manifest
	manifest := &models.Manifest{}
	if err := json.Unmarshal(fbytes, &manifest); err != nil {
		return fmt.Errorf("error parsing manifest: %w", err)
	}

	pp.Println("LoadPtypeWithFolder/3")

	// Get project type from manifest
	ptype := manifest.Slug

	// Create loaded definition
	ldef := LoadedDef{
		defType: defTypeLuaFolder,
		file:    filePath,
	}

	// Create base path for project
	basePath := "/z/p/" + ptype
	serveFolder := manifest.ServeFolder
	serverFile := manifest.ServerFile

	if serverFile == "" {
		serverFile = "server.lua"
	}

	linkPattern := manifest.LinkPattern
	if linkPattern == "" {
		linkPattern = fmt.Sprintf("%s?pid={PID}", basePath)
	}

	pp.Println("LoadPtypeWithFolder/4")

	root, err := os.OpenRoot(filePath)
	if err != nil {
		return err
	}

	pp.Println("LoadPtypeWithFolder/5")

	svrFile, err := root.OpenFile(serverFile, os.O_RDONLY, 0766)
	if err != nil {
		pp.Println("serverFile", serverFile)
		pp.Println("LoadPtypeWithFolder/5.1", err)
		return err
	}

	defer svrFile.Close()

	svrFileBytes, err := io.ReadAll(svrFile)
	if err != nil {
		return err
	}

	pp.Println("LoadPtypeWithFolder/6")

	lexecutor := luaz.New(luaz.Options{
		BuilderOpts: xproject.BuilderOption{
			App:    e.opts.App,
			Logger: e.opts.Logger,
		},
		Code:          string(svrFileBytes),
		WorkingFolder: path.Join(e.opts.BasePath, "wd"),
	})

	pp.Println("LoadPtypeWithFolder/7")

	// Create project definition
	def := &xproject.Defination{
		Name:          manifest.Name,
		Slug:          manifest.Slug,
		Info:          manifest.Info,
		OnPageRequest: ServeFolderContentsWithPrefix(root, basePath, serveFolder),
		OnProjectRequest: func(ctx *gin.Context) {
			pp.Println("@onProjectRequest", ctx.Request.URL.Path)
			lexecutor.Handle(ctx)
		},
		OnClose: func() error {

			return root.Close()
		},
		LinkPattern: linkPattern,
	}

	pp.Println("LoadPtypeWithFolder/8")

	ldef.def = def

	e.pLock.Lock()
	e.projects[ptype] = &ldef
	e.pLock.Unlock()

	return nil
}

func (e *Engine) LoadPtypeWithZip(filePath string) error {

	manifest := &models.Manifest{}
	err := ReadManifestFromZip(filePath, manifest)
	if err != nil {
		fmt.Println(err)
		return err
	}

	ptype := manifest.Slug
	ldef := LoadedDef{
		defType: defTypeLuaZip,
		file:    filePath,
	}

	r, err := zip.OpenReader(filePath)
	if err != nil {
		return fmt.Errorf("error opening zip file: %w", err)
	}

	serveFolder := manifest.ServeFolder

	basePath := "/z/p/" + ptype
	linkPattern := manifest.LinkPattern
	if linkPattern == "" {
		linkPattern = fmt.Sprintf("%s?pid={PID}", basePath)
	}

	serverFile := manifest.ServerFile

	if serverFile == "" {
		serverFile = "server.lua"
	}

	svrFile, err := r.Open(serverFile)
	if err != nil {
		return err
	}

	defer svrFile.Close()

	svrFileBytes, err := io.ReadAll(svrFile)
	if err != nil {
		return err
	}

	lexecutor := luaz.New(luaz.Options{
		BuilderOpts: xproject.BuilderOption{
			App:    e.opts.App,
			Logger: e.opts.Logger,
		},
		Code:          string(svrFileBytes),
		WorkingFolder: path.Join(e.opts.BasePath, "wd"),
	})

	def := &xproject.Defination{
		Name:          manifest.Name,
		Slug:          manifest.Slug,
		Info:          manifest.Info,
		OnPageRequest: ServeZipContentsWithPrefix(r, basePath, serveFolder),
		OnProjectRequest: func(ctx *gin.Context) {
			pp.Println("@onProjectRequest", ctx.Request.URL.Path)
			lexecutor.Handle(ctx)
		},
		OnClose: func() error {
			return r.Close()
		},
		LinkPattern: linkPattern,
	}

	ldef.def = def

	e.pLock.Lock()
	e.projects[ptype] = &ldef
	e.pLock.Unlock()

	return nil

}
