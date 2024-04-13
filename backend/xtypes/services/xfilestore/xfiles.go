package xfilestore

import (
	"context"
	"io"
)

type BlobInfo struct {
	Name         string `json:"name,omitempty"`
	IsDir        bool   `json:"is_dir,omitempty"`
	Size         int    `json:"size,omitempty"`
	LastModified string `json:"last_modified,omitempty"`
}

type FilesService interface {
	ProjectFiles(pid int64) FileStore
	UserFiles(uid int64) FileStore
}

type FileStore interface {
	ListFolder(ctx context.Context, fpath string) ([]*BlobInfo, error)
	NewFolder(ctx context.Context, fpath, name string) error
	DeleteFolder(ctx context.Context, fpath string) error
	RenameFolder(ctx context.Context, fpath, newname string) error
	TreeFolder(ctx context.Context, fpath string) ([]*BlobInfo, error)
	GetFile(ctx context.Context, fpath string) (FData, error)
	RenameFile(ctx context.Context, fpath, name, newname string) error
	DuplicateFile(ctx context.Context, fpath, name, newname string) error
	MoveFile(ctx context.Context, fpath, newfpath string) error
	NewFile(ctx context.Context, fpath, name string, data FData) error
	UpdateFile(ctx context.Context, fpath, name string, data FData) error
	DeleteFile(ctx context.Context, fpath, name string) error
	CompressFolder(ctx context.Context, fpath string) (FData, error)
	CompressFiles(ctx context.Context, fpath string, files []string) (FData, error)
}

type FData interface {
	AsBytes() ([]byte, error)
	AsReader() (io.Reader, error)
	Close() error
}
