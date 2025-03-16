package engine

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/blue-monads/turnix/backend/xtypes/xproject"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
)

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
			pp.Println("@skip_path", path)
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

func (e *Engine) LoadPtypeWithZip(filePath string) error {
	//	file := fmt.Sprintf("%s/%s.zip", e.installPath, ptype)

	manifest := &Manifest{}
	err := ReadManifestFromZip(filePath, manifest)
	if err != nil {
		fmt.Println(err)
		return err
	}

	ptype := manifest.Slug

	def := &xproject.Defination{
		Name: manifest.Name,
		Slug: manifest.Slug,
		Info: manifest.Info,
		OnPageRequest: func(ctx *gin.Context) {
			pp.Println("@OnPageRequest", ctx.Request.URL)

		},
		OnProjectRequest: func(ctx *gin.Context) {
			pp.Println("@OnProjectRequest", ctx.Request.URL)

		},
		LinkPattern: "/z/projects/" + ptype,
	}

	e.pLock.Lock()
	e.projects[ptype] = def
	e.pLock.Unlock()

	return nil

}

const manifestFileName = "manifest.json"

func ReadManifestFromZip(zipFilePath string, target any) error {
	r, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return fmt.Errorf("error opening zip file: %w", err)
	}
	defer r.Close()

	for _, f := range r.File {
		if f.Name == manifestFileName {
			rc, err := f.Open()
			if err != nil {
				return fmt.Errorf("error opening '%s' in zip: %w", manifestFileName, err)
			}
			defer rc.Close()

			content, err := io.ReadAll(rc)
			if err != nil {
				return fmt.Errorf("error reading content of '%s': %w", manifestFileName, err)
			}

			err = json.Unmarshal(content, target)
			if err != nil {
				return fmt.Errorf("error unmarshalling JSON from '%s': %w", manifestFileName, err)
			}

			return nil
		}
	}

	return fmt.Errorf("'%s' not found in zip archive", manifestFileName)
}

func DownloadFile(fileURL, outputPath string) error {
	resp, err := http.Get(fileURL)
	if err != nil {
		return fmt.Errorf("error fetching file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK status code: %d", resp.StatusCode)
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer outputFile.Close()

	buffer := make([]byte, 4096)
	_, err = io.CopyBuffer(outputFile, resp.Body, buffer)
	if err != nil {
		return fmt.Errorf("error streaming and writing file: %w", err)
	}

	return nil
}
