package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
	"os/exec"

	"github.com/JustinTimperio/osinfo"

	"github.com/kbinani/screenshot"
	"github.com/mackerelio/go-osstat/memory"
)

var (
	MaxBytes = int64(1024 * 1024 * 5)
)

func handlePing(ctx *WHContext) (any, error) {
	return "pong", nil
}

func handleFsListDir(ctx *WHContext) (any, error) {
	path := ctx.GetAsString("path")

	if path == "" {
		path = "/"
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to list directory: %v", err)
	}

	result := make([]string, len(entries))
	for i, entry := range entries {
		result[i] = entry.Name()
	}
	return result, nil
}

func handleFsReadFile(ctx *WHContext) (any, error) {
	path := ctx.GetAsString("path")

	// Check file size first
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %v", err)
	}

	// Limit to 1MB (1048576 bytes)
	if fileInfo.Size() > (MaxBytes) {
		return nil, fmt.Errorf("file is too large (max %d bytes): %d bytes", MaxBytes, fileInfo.Size())
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}
	return content, nil
}

func handleFsWriteFile(ctx *WHContext) (any, error) {
	path := ctx.GetAsString("path")
	data := ctx.GetAsString("data")
	isBase64 := ctx.GetAsBool("base64")

	var fileData []byte
	var err error

	if isBase64 {
		fileData, err = base64.StdEncoding.DecodeString(data)
		if err != nil {
			return nil, fmt.Errorf("failed to decode base64 data: %v", err)
		}
	} else {
		fileData = []byte(data)
	}

	err = os.WriteFile(path, fileData, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to write file: %v", err)
	}
	return nil, nil
}

func handleFsReadRange(ctx *WHContext) (any, error) {
	path := ctx.GetAsString("path")
	start := ctx.GetAsInt("start")
	end := ctx.GetAsInt("end")

	// Open the file
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Get file size
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %v", err)
	}
	fileSize := fileInfo.Size()

	// Validate and adjust start and end positions
	if start < 0 {
		start = 0
	}
	if end <= 0 || end > int(fileSize) {
		end = int(fileSize)
	}

	// Ensure start is not beyond end
	if start > end {
		return nil, fmt.Errorf("invalid range: start (%d) is greater than end (%d)", start, end)
	}

	// Calculate range length
	rangeLength := end - start

	if int64(rangeLength) > MaxBytes {
		return nil, fmt.Errorf("file range is too large (max %d bytes): %d bytes", MaxBytes, rangeLength)
	}

	// Create buffer and read specific range
	buffer := make([]byte, rangeLength)
	_, err = file.Seek(int64(start), 0)
	if err != nil {
		return nil, fmt.Errorf("failed to seek to start position: %v", err)
	}

	bytesRead, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("failed to read file range: %v", err)
	}

	return buffer[:bytesRead], nil
}

func handleFsRemove(ctx *WHContext) (any, error) {
	path := ctx.GetAsString("path")
	err := os.Remove(path)
	if err != nil {
		return nil, fmt.Errorf("failed to remove: %v", err)
	}
	return nil, nil
}

func handleFsMkdir(ctx *WHContext) (any, error) {
	path := ctx.GetAsString("path")
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to create directory: %v", err)
	}
	return nil, nil
}

func handleFsRename(ctx *WHContext) (any, error) {
	oldPath := ctx.GetAsString("oldPath")
	newPath := ctx.GetAsString("newPath")
	err := os.Rename(oldPath, newPath)
	if err != nil {
		return nil, fmt.Errorf("failed to rename: %v", err)
	}
	return nil, nil
}

func SystemExec(ctx *WHContext) (any, error) {
	command := ctx.GetAsString("command")
	output, err := exec.Command("sh", "-c", command).Output()
	if err != nil {
		return nil, fmt.Errorf("command execution failed: %v", err)
	}

	return string(output), nil
}

type SystemInfoResponse struct {
	OsInfo   osinfo.Release `json:"os_info"`
	HostName string         `json:"hostname"`
	Memory   *memory.Stats  `json:"memory"`
}

func handleSystemInfo(ctx *WHContext) (any, error) {

	osinfo := osinfo.GetVersion()
	hostname, _ := os.Hostname()

	memory, err := memory.Get()
	if err != nil {
		return nil, fmt.Errorf("failed to get memory info: %v", err)
	}

	r := SystemInfoResponse{
		OsInfo:   osinfo,
		HostName: hostname,
		Memory:   memory,
	}

	return r, nil
}

func handleScreenShot(ctx *WHContext) (any, error) {

	displayIndex := ctx.GetAsInt("displayIndex")
	bounds := screenshot.GetDisplayBounds(displayIndex)

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return nil, fmt.Errorf("failed to capture screenshot: %v", err)
	}

	buf := bytes.NewBuffer(nil)
	err = png.Encode(buf, img)
	if err != nil {
		return nil, fmt.Errorf("failed to encode screenshot: %v", err)
	}

	// write to file

	finalBytes := buf.Bytes()

	// err = os.WriteFile("screenshot.png", finalBytes, 0644)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to write screenshot to file: %v", err)
	// }

	return finalBytes, nil
}

func handleListDisplays(ctx *WHContext) (any, error) {

	n := screenshot.NumActiveDisplays()

	displays := make([]image.Rectangle, n)

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		displays[i] = bounds
	}

	return displays, nil
}

func init() {
	RegisterHandler("ping", handlePing)
	RegisterHandler("fs.listdir", handleFsListDir)
	RegisterHandler("fs.readfile", handleFsReadFile)
	RegisterHandler("fs.writefile", handleFsWriteFile)
	RegisterHandler("fs.readrange", handleFsReadRange)
	RegisterHandler("fs.remove", handleFsRemove)
	RegisterHandler("fs.mkdir", handleFsMkdir)
	RegisterHandler("fs.rename", handleFsRename)
	RegisterHandler("system.exec", SystemExec)
	RegisterHandler("system.info", handleSystemInfo)
	RegisterHandler("system.listDisplays", handleListDisplays)
	RegisterHandler("system.screenshot", handleScreenShot)
	RegisterHandler("service.hello", handleServiceHello)
	RegisterHandler("service.shell", handleServiceShell)
}
