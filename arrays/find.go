package arrays

import "reflect"

type find struct{}

var Find = find{}

func (f find) Interface(list interface{}, fn interface{}) (interface{}, bool) {
	reflectValue := reflect.ValueOf(list)
	typ := reflectValue.Type()
	if reflectValue.IsNil() {
		return reflect.Zero(typ.Elem()).Interface(), false
	}

	length := reflectValue.Len()
	for i := 0; i < length; i++ {
		target := reflectValue.Index(i)
		if reflect.ValueOf(fn).Call([]reflect.Value{target})[0].Bool() {
			return target.Interface(), true
		}
	}
	return reflect.Zero(typ.Elem()).Interface(), false
}
