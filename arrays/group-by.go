package arrays

import "reflect"

type groupByImpl struct{}

var GroupBy = groupByImpl{}

func (m groupByImpl) Interface(list interface{}, fn interface{}) interface{} {
	reflectValue := reflect.ValueOf(list)
	funcValue := reflect.ValueOf(fn)
	funcReturnType := funcValue.Type().Out(0)
	if reflectValue.IsNil() {
		return reflect.MakeMap(reflect.MapOf(funcReturnType, reflect.TypeOf(list))).Interface()
	}

	length := reflectValue.Len()
	result := reflect.MakeMap(reflect.MapOf(funcReturnType, reflect.TypeOf(list)))
	for i := 0; i < length; i++ {
		mapKey := reflect.ValueOf(fn).Call([]reflect.Value{reflectValue.Index(i)})[0]
		mapValue := result.MapIndex(mapKey)
		if !mapValue.IsValid() {
			result.SetMapIndex(mapKey, reflect.MakeSlice(reflect.TypeOf(list), 0, 0))
			mapValue = result.MapIndex(mapKey)
		}
		result.SetMapIndex(mapKey, reflect.Append(mapValue, reflectValue.Index(i)))
	}
	for _, x := range result.MapKeys() {
		mapValue := result.MapIndex(x)
		result.SetMapIndex(x, mapValue.Slice3(0, mapValue.Len(), mapValue.Len()))
	}
	return result.Interface()
}
