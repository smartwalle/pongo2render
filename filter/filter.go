package filter

import (
	"github.com/flosch/pongo2"
	"github.com/smartwalle/container"
)

func init() {
	pongo2.RegisterFilter("GetValueWithMap", GetValueWithMap)
}

func GetValueWithMap(in, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
	var v = container.GetValueWithMap(in.Interface(), param.Interface())
	out = pongo2.AsValue(v)
	return out, err
}