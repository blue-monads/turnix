package engine

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

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

type Manifest struct {
	Name     string         `json:"name"`
	Slug     string         `json:"slug"`
	Info     string         `json:"info"`
	Type     string         `json:"type"`
	Format   string         `json:"format"`
	Tags     []string       `json:"tags"`
	Routes   []Route        `json:"routes"`
	Services map[string]any `json:"services"`
}

type Route struct {
	Name     string            `json:"name"`
	Type     string            `json:"type"` // authed_http, http, ws
	Method   string            `json:"method"`
	Path     string            `json:"path"`
	Handlers map[string]string `json:"handlers"`
	Options  map[string]any    `json:"options"`
}

func (e *Engine) load() {

	filepath.WalkDir(e.installPath, func(path string, d fs.DirEntry, err error) error {
		pp.Println("@curr_path", path)

		if d.IsDir() {
			err = e.LoadPtypeWithFolder(path)
			if err != nil {
				pp.Println("@skipping_folder", path, err)
				return nil
			}

			return nil
		}

		if strings.Contains(path, "__downloading_") {
			pp.Println("@skipping_downloading", path)
			return nil
		}

		if strings.HasSuffix(path, ".zip") {
			err = e.LoadPtypeWithZip(path)
			if err != nil {
				pp.Println("@skipping_zip", path, err)
				return nil
			}
		}

		return nil

	})

}

func (e *Engine) LoadPtypeWithFolder(filePath string) error {
	// Read the manifest file
	manifestFile := path.Join(filePath, manifestFileName)
	fbytes, err := os.ReadFile(manifestFile)
	if err != nil {
		return err
	}

	// Parse manifest
	manifest := &Manifest{}
	if err := json.Unmarshal(fbytes, &manifest); err != nil {
		return fmt.Errorf("error parsing manifest: %w", err)
	}

	// Get project type from manifest
	ptype := manifest.Slug

	// Create loaded definition
	ldef := LoadedDef{
		defType: defTypeLuaFolder,
		file:    filePath,
	}

	// Create base path for project
	basePath := "/z/p/" + ptype

	root, err := os.OpenRoot(filePath)
	if err != nil {
		return err
	}

	// Create project definition
	def := &xproject.Defination{
		Name:          manifest.Name,
		Slug:          manifest.Slug,
		Info:          manifest.Info,
		OnPageRequest: ServeFolderContentsWithPrefix(root, basePath),
		OnProjectRequest: func(ctx *gin.Context) {

		},
		OnClose: func() error {
			return root.Close()
		},
		LinkPattern: basePath,
	}

	// Assign definition to loaded definition
	ldef.def = def

	// Register the project in the engine
	e.pLock.Lock()
	e.projects[ptype] = &ldef
	e.pLock.Unlock()

	return nil
}

func (e *Engine) LoadPtypeWithZip(filePath string) error {

	manifest := &Manifest{}
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

	basePath := "/z/p/" + ptype

	def := &xproject.Defination{
		Name:          manifest.Name,
		Slug:          manifest.Slug,
		Info:          manifest.Info,
		OnPageRequest: ServeZipContentsWithPrefix(r, basePath),
		OnProjectRequest: func(ctx *gin.Context) {
			// fixme => impl
		},
		OnClose: func() error {
			return r.Close()
		},
		LinkPattern: basePath,
	}

	ldef.def = def

	e.pLock.Lock()
	e.projects[ptype] = &ldef
	e.pLock.Unlock()

	return nil

}
