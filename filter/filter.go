package filter

import (
	"reflect"
	"github.com/flosch/pongo2"
	"strings"
)

func init() {
	pongo2.RegisterFilter("ValueWithMap", ValueWithMap)
	pongo2.RegisterFilter("HasPrefix", HasPrefix)
	pongo2.RegisterFilter("HasSuffix", HasSuffix)
	pongo2.RegisterFilter("EqualToString", EqualToString)
	pongo2.RegisterFilter("CompareString", CompareString)
}

////////////////////////////////////////////////////////////////////////////////
func ValueWithMap(in, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
	var v = valueWithMap(in.Interface(), param.Interface())
	out = pongo2.AsValue(v)
	return out, err
}

func valueWithMap(source, key interface{}) interface{} {
	if source == nil {
		return nil
	}

	var sourceValue = reflect.ValueOf(source)
	if sourceValue.IsNil() {
		return nil
	}

	switch sourceValue.Kind() {
	case reflect.Map:
		var targetValue = reflect.ValueOf(key)

		if targetValue.IsValid() {
			return sourceValue.MapIndex(targetValue).Interface()
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
func HasPrefix(in, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	return pongo2.AsValue(strings.HasPrefix(in.String(), param.String())), nil
}

func HasSuffix(in, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	return pongo2.AsValue(strings.HasSuffix(in.String(), param.String())), nil
}

func EqualToString(in, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	return pongo2.AsValue(in.String() == param.String()), nil
}

func CompareString(in, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	return pongo2.AsValue(strings.Compare(in.String(), param.String())), nil
}