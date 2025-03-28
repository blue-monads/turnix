package engine

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
)

func ServeFolderContentsWithPrefix(root *os.Root, pathPrefixToRemove string) gin.HandlerFunc {
	// Create a root filesystem from the base folder to prevent path traversal

	return func(c *gin.Context) {
		requestPath := c.Request.URL.Path
		if pathPrefixToRemove != "" && strings.HasPrefix(requestPath, pathPrefixToRemove) {
			requestPath = strings.TrimPrefix(requestPath, pathPrefixToRemove)
		}

		requestPath = strings.TrimPrefix(requestPath, "/")

		// Open the file using the root filesystem to prevent path traversal
		file, err := root.Open(requestPath)
		if err != nil {
			if os.IsNotExist(err) {
				c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
				log.Printf("Error opening file: %v", err)
			}
			return
		}
		defer file.Close()

		// Get file info to check if it's a directory
		stat, err := file.Stat()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get file info"})
			log.Printf("Error getting file info: %v", err)
			return
		}

		// Don't serve directories
		if stat.IsDir() {
			c.JSON(http.StatusForbidden, gin.H{"error": "Cannot serve directories"})
			return
		}

		// Set content type based on file extension
		contentType := mime.TypeByExtension(filepath.Ext(requestPath))
		if contentType != "" {
			c.Writer.Header().Set("Content-Type", contentType)
		}

		// Set content length
		c.Writer.Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size()))

		pp.Println("@serveFolderContentsWithPrefix/final_boss", requestPath)

		// Copy the file contents to the response
		_, err = io.Copy(c.Writer, file)
		if err != nil {
			log.Printf("Error serving file: %v", err)
		}
	}
}
func ServeZipContentsWithPrefix(zipfile *zip.ReadCloser, pathPrefixToRemove string) gin.HandlerFunc {

	fileMap := make(map[string]*zip.File)
	for _, f := range zipfile.File {
		zipPath := strings.TrimPrefix(f.Name, "/")
		fileMap[zipPath] = f
	}

	return func(c *gin.Context) {

		requestPath := c.Request.URL.Path

		pp.Println("@zipserve/requestPath", requestPath)

		if requestPath != "" && strings.HasPrefix(requestPath, pathPrefixToRemove) {
			requestPath = strings.TrimPrefix(requestPath, pathPrefixToRemove)
		}

		requestPath = strings.TrimPrefix(requestPath, "/")

		if requestPath == "" {
			requestPath = "index.html"
		}

		file, exists := fileMap[requestPath]
		if !exists {
			file, exists = fileMap[fmt.Sprintf("%s.html", requestPath)]
			if !exists {
				c.AbortWithError(http.StatusNotFound, errors.New("not found"))
				return
			}
		}

		fileReader, err := file.Open()
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		defer fileReader.Close()

		contentType := mime.TypeByExtension(filepath.Ext(file.Name))
		if contentType != "" {
			c.Writer.Header().Set("Content-Type", contentType)
		}

		c.Request.Header.Set("Content-Length", fmt.Sprintf("%d", file.UncompressedSize64))

		_, err = io.Copy(c.Writer, fileReader)
		if err != nil {
			log.Printf("Error serving zip file: %v\n", err)
		}
	}
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
