package chunk

import (
	"reflect"
)

func Interface(list interface{}, chunkSize int) interface{} {
	if list == nil {
		return []interface{}{}
	}
	reflectValue := reflect.ValueOf(list)
	typ := reflectValue.Type()
	length := reflectValue.Len()

	result := reflect.MakeSlice(reflect.SliceOf(typ), 0, 0)
	for i := 0; i < length; {
		temp := reflect.MakeSlice(typ, 0, 0)
		for j := 0; i+j < length && j < chunkSize; i, j = i+1, j+1 {
			temp = reflect.Append(temp, reflectValue.Index(i))
		}
		result = reflect.Append(result, temp)
	}
	return result.Interface()
}

func Int(list []int, chunkSize int) [][]int {
	return Interface(list, chunkSize).([][]int)
}

func String(list []string, chunkSize int) [][]string {
	return Interface(list, chunkSize).([][]string)
}
