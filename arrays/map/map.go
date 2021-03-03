package filter

import "reflect"

func Interface(list interface{}, f interface{}) interface{} {
	reflectValue := reflect.ValueOf(list)
	funcValue := reflect.ValueOf(f)
	funcReturnType := funcValue.Type().Out(0)
	if reflectValue.IsNil() {
		return reflect.MakeSlice(reflect.SliceOf(funcReturnType), 0, 0).Interface()
	}

	length := reflectValue.Len()
	result := reflect.MakeSlice(reflect.SliceOf(funcReturnType), 0, length)
	for i := 0; i < length; i++ {
		result = reflect.Append(result, reflect.ValueOf(f).Call([]reflect.Value{reflectValue.Index(i)})[0])
	}
	return result.Slice3(0, result.Len(), result.Len()).Interface()
}
