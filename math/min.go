package math

import (
	"math"
	"reflect"
)

type min struct{}

var Min = min{}

func (m min) Interface(value interface{}, valueList interface{}) interface{} {
	result := ToFloat64(value)
	valueListReflectValue := reflect.ValueOf(valueList)
	length := valueListReflectValue.Len()
	for i := 0; i < length; i++ {
		result = math.Min(result, ToFloat64(valueListReflectValue.Index(i).Interface()))
	}

	return reflect.ValueOf(result).Convert(reflect.ValueOf(value).Type()).Interface()
}

func (m min) Float64(value float64, valueList ...float64) float64 {
	return m.Interface(value, valueList).(float64)
}

func (m min) Int(value int, valueList ...int) int {
	return m.Interface(value, valueList).(int)
}

func (m min) UInt(value uint, valueList ...uint) uint {
	return m.Interface(value, valueList).(uint)
}
