package lazydata

import (
	"encoding/json"

	"github.com/temphia/temphia/code/backend/xtypes"
)

var (
	_ xtypes.LazyData = (*AnyData)(nil)
)

type AnyData struct {
	inner any
}

func NewAnyData(data any) *AnyData {
	return &AnyData{
		inner: data,
	}
}

func (a *AnyData) AsJsonBytes() ([]byte, error) { return json.Marshal(&a.inner) }
func (a *AnyData) AsObject(target any) error {
	out, err := json.Marshal(&a.inner)
	if err != nil {
		return err
	}

	return json.Unmarshal(out, target)
}
func (a *AnyData) IsJsonBytes() bool { return false }
func (a *AnyData) IsObject() bool    { return false }
func (a *AnyData) Inner() any        { return a.inner }
