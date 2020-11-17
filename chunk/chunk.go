package chunk

import (
	"reflect"

	"github.com/ryo-kagawa/go-utils/conditional"
)

func min(value1, value2 int) int {
	return conditional.Int(value1 < value2, value1, value2)
}

func Interface(list interface{}, chunkSize int) interface{} {
	if list == nil {
		return []interface{}{}
	}
	reflectValue := reflect.ValueOf(list)
	typ := reflectValue.Type()
	length := reflectValue.Len()

	result := reflect.MakeSlice(reflect.SliceOf(typ), 0, length/chunkSize+min(length%chunkSize, 1))
	for i := 0; i < length; i += chunkSize {
		result = reflect.Append(result, reflectValue.Slice(i, min(i+chunkSize, length)))
	}
	return result.Interface()
}

func Int(list []int, chunkSize int) [][]int {
	return Interface(list, chunkSize).([][]int)
}

func String(list []string, chunkSize int) [][]string {
	return Interface(list, chunkSize).([][]string)
}
