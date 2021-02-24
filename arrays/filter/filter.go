package filter

import "reflect"

func Interface(list interface{}, f interface{}) interface{} {
	if list == nil {
		return list
	}

	reflectValue := reflect.ValueOf(list)
	typ := reflectValue.Type()
	length := reflectValue.Len()

	result := reflect.MakeSlice(typ, 0, 0)
	for i := 0; i < length; i++ {
		target := reflectValue.Index(i)
		if reflect.ValueOf(f).Call([]reflect.Value{target})[0].Bool() {
			result = reflect.Append(result, target)
		}
	}
	return result.Interface()
}

func Int(list []int, f func(v int) bool) []int {
	return Interface(list, f).([]int)
}

func String(list []string, f func(v string) bool) []string {
	return Interface(list, f).([]string)
}
