package files

import (
	"fmt"

	"github.com/bornjre/trunis/backend/services/filestore/localfs"
	"github.com/bornjre/trunis/backend/xtypes/services/xfilestore"
)

type FilesService struct {
	baseFile string
}

func New(baseFile string) *FilesService {
	return &FilesService{baseFile: baseFile}
}

func (f *FilesService) ProjectFiles(pid int64) xfilestore.FileStore {
	return localfs.New(fmt.Sprintf("%s/projects/%d", f.baseFile, pid))
}

func (f *FilesService) UserFiles(uid int64) xfilestore.FileStore {
	return localfs.New(fmt.Sprintf("%s/projects/%d", f.baseFile, uid))
}
