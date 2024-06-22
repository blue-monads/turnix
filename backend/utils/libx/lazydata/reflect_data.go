package lazydata

import (
	"encoding/json"
	"reflect"
)

type ReflectData struct {
	val reflect.Value
}

func NewReflectData(val reflect.Value) *ReflectData {
	return &ReflectData{
		val: val,
	}
}

func (a *ReflectData) AsJsonBytes() ([]byte, error) {
	return json.Marshal(a.val.Interface())
}

func (a *ReflectData) AsObject(target any) error {
	bytes, err := a.AsJsonBytes()
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, target)
}
func (a *ReflectData) IsJsonBytes() bool { return false }
func (a *ReflectData) IsObject() bool    { return false }
func (a *ReflectData) Inner() any        { return a.val }
