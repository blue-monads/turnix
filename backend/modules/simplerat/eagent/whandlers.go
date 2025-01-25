package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"

	"github.com/JustinTimperio/osinfo"

	"github.com/mackerelio/go-osstat/memory"
)

var (
	MaxBytes = int64(1024 * 1024 * 5)
)

type PingResponse struct {
	OsInfo   osinfo.Release `json:"os_info"`
	Message  string         `json:"message"`
	HostName string         `json:"hostname"`
	Memory   *memory.Stats  `json:"memory"`
}

func handlePing(ctx *WHContext) (any, error) {
	osinfo := osinfo.GetVersion()
	hostname, _ := os.Hostname()

	memory, err := memory.Get()
	if err != nil {
		return nil, fmt.Errorf("failed to get memory info: %v", err)
	}

	r := PingResponse{
		OsInfo:   osinfo,
		Message:  "pong",
		HostName: hostname,
		Memory:   memory,
	}

	return r, nil
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

// Don't forget to register these handlers in your init or main function
func init() {
	RegisterHandler("ping", handlePing)
	RegisterHandler("fs.listdir", handleFsListDir)
	RegisterHandler("fs.readfile", handleFsReadFile)
	RegisterHandler("fs.writefile", handleFsWriteFile)
	RegisterHandler("fs.remove", handleFsRemove)
	RegisterHandler("fs.mkdir", handleFsMkdir)
	RegisterHandler("fs.rename", handleFsRename)
	RegisterHandler("system.exec", SystemExec)
}
