package lazydata

import (
	"encoding/json"

	"github.com/temphia/temphia/code/backend/xtypes"
)

var (
	_ xtypes.LazyData = (*JsonData)(nil)
)

type JsonData struct {
	inner []byte
}

func NewJsonData(data []byte) *JsonData {
	return &JsonData{
		inner: data,
	}
}

func (j *JsonData) AsJsonBytes() ([]byte, error) { return j.inner, nil }
func (j *JsonData) AsObject(target any) error    { return json.Unmarshal(j.inner, target) }

func (j *JsonData) IsJsonBytes() bool { return true }
func (j *JsonData) IsObject() bool    { return false }
func (j *JsonData) Inner() any        { return nil }
