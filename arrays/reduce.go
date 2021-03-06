package arrays

import "reflect"

type reduce struct{}

var Reduce = reduce{}

func (r reduce) Interface(list interface{}, fn interface{}) interface{} {
	reflectValue := reflect.ValueOf(list)
	funcValue := reflect.ValueOf(fn)
	funcReturnType := funcValue.Type().Out(0)
	if reflectValue.IsNil() {
		return reflect.Zero(funcReturnType).Interface()
	}

	length := reflectValue.Len()
	funcValueLength := funcValue.Type().NumIn()
	accumulator := reflect.New(funcReturnType).Elem()
	for i := 0; i < length; i++ {
		value := []reflect.Value{}
		value = append(value, accumulator)
		value = append(value, reflectValue.Index(i))
		if funcValueLength >= 3 {
			value = append(value, reflect.ValueOf(i))
		}
		if funcValueLength >= 4 {
			value = append(value, reflectValue)
		}
		accumulator = reflect.ValueOf(fn).Call(value)[0]
	}
	return accumulator.Interface()
}
