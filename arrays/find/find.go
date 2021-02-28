package find

import "reflect"

func Interface(list interface{}, f interface{}) (interface{}, bool) {
	reflectValue := reflect.ValueOf(list)
	typ := reflectValue.Type()
	if reflectValue.IsNil() {
		return reflect.Zero(typ.Elem()).Interface(), false
	}

	length := reflectValue.Len()
	for i := 0; i < length; i++ {
		target := reflectValue.Index(i)
		if reflect.ValueOf(f).Call([]reflect.Value{target})[0].Bool() {
			return target.Interface(), true
		}
	}
	return reflect.Zero(typ.Elem()).Interface(), false
}

func Int(list []int, f func(v int) bool) (int, bool) {
	x, ok := Interface(list, f)
	return x.(int), ok
}

func String(list []string, f func(v string) bool) (string, bool) {
	x, ok := Interface(list, f)
	return x.(string), ok
}
