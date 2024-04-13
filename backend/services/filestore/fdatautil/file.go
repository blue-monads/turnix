package fdatautil

import (
	"io"
	"os"

	"github.com/bornjre/trunis/backend/xtypes/services/xfilestore"
)

var _ xfilestore.FData = (*FileFdata)(nil)

type FileFdata struct {
	fpath         string
	file          *os.File
	deleteOnClose bool
}

func NewFromFile(fpath string, delonclose bool) *FileFdata {
	return &FileFdata{
		fpath:         fpath,
		file:          nil,
		deleteOnClose: delonclose,
	}
}

func (f *FileFdata) AsBytes() ([]byte, error) {
	return os.ReadFile(f.fpath)
}

func (f *FileFdata) AsReader() (io.Reader, error) {

	file, err := os.Open(f.fpath)
	if err != nil {
		return nil, err
	}

	f.file = file

	return file, nil
}

func (f *FileFdata) Close() error {
	if f.file != nil {
		f.file.Close()
	}

	if f.deleteOnClose {
		return os.Remove(f.fpath)
	}

	return nil

}

func (f *FileFdata) InnerFile() string {
	return f.fpath
}
