package chunk

import (
	"reflect"

	"github.com/ryo-kagawa/go-utils/math/min"
)

func Interface(list interface{}, chunkSize int) interface{} {
	reflectValue := reflect.ValueOf(list)
	typ := reflectValue.Type()
	if reflectValue.IsNil() {
		return reflect.MakeSlice(reflect.SliceOf(typ), 0, 0).Interface()
	}

	length := reflectValue.Len()
	result := reflect.MakeSlice(reflect.SliceOf(typ), 0, length/chunkSize+min.Int(length%chunkSize, 1))
	for i := 0; i < length; i += chunkSize {
		rangeLast := min.Int(i+chunkSize, length)
		result = reflect.Append(result, reflectValue.Slice3(i, rangeLast, rangeLast))
	}
	return result.Interface()
}

func Int(list []int, chunkSize int) [][]int {
	return Interface(list, chunkSize).([][]int)
}

func String(list []string, chunkSize int) [][]string {
	return Interface(list, chunkSize).([][]string)
}
