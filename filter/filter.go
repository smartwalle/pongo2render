package filter

import (
	"reflect"
	"github.com/flosch/pongo2"
)

func init() {
	pongo2.RegisterFilter("GetValueWithMap", GetValueWithMap)
}

func GetValueWithMap(in, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
	var v = getValueWithMap(in.Interface(), param.Interface())
	out = pongo2.AsValue(v)
	return out, err
}

func getValueWithMap(source, key interface{}) interface{} {
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