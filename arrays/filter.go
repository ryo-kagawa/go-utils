package arrays

import "reflect"

type filter struct{}

var Filter = filter{}

func (f filter) Interface(list interface{}, fn interface{}) interface{} {
	reflectValue := reflect.ValueOf(list)
	typ := reflectValue.Type()
	if reflectValue.IsNil() {
		return reflect.MakeSlice(typ, 0, 0).Interface()
	}

	length := reflectValue.Len()
	result := reflect.MakeSlice(typ, 0, length)
	for i := 0; i < length; i++ {
		target := reflectValue.Index(i)
		if reflect.ValueOf(fn).Call([]reflect.Value{target})[0].Bool() {
			result = reflect.Append(result, target)
		}
	}
	return result.Slice3(0, result.Len(), result.Len()).Interface()
}
