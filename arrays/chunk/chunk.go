package chunk

import (
	"reflect"

	"github.com/ryo-kagawa/go-utils/math"
)

func Interface(list interface{}, chunkSize int) interface{} {
	reflectValue := reflect.ValueOf(list)
	typ := reflectValue.Type()
	if reflectValue.IsNil() {
		return reflect.MakeSlice(reflect.SliceOf(typ), 0, 0).Interface()
	}

	length := reflectValue.Len()
	result := reflect.MakeSlice(reflect.SliceOf(typ), 0, length/chunkSize+math.Min.Int(length%chunkSize, 1))
	for i := 0; i < length; i += chunkSize {
		rangeLast := math.Min.Int(i+chunkSize, length)
		result = reflect.Append(result, reflectValue.Slice3(i, rangeLast, rangeLast))
	}
	return result.Interface()
}
