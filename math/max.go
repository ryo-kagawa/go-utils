package math

import (
	"math"
	"reflect"
)

type max struct{}

var Max = max{}

func (m max) Interface(value interface{}, valueList interface{}) interface{} {
	result := ToFloat64(value)
	valueListReflectValue := reflect.ValueOf(valueList)
	length := valueListReflectValue.Len()
	for i := 0; i < length; i++ {
		result = math.Max(result, ToFloat64(valueListReflectValue.Index(i).Interface()))
	}

	return reflect.ValueOf(result).Convert(reflect.ValueOf(value).Type()).Interface()
}

func (m max) Float64(value float64, valueList ...float64) float64 {
	return m.Interface(value, valueList).(float64)
}

func (m max) Int(value int, valueList ...int) int {
	return m.Interface(value, valueList).(int)
}

func (m max) UInt(value uint, valueList ...uint) uint {
	return m.Interface(value, valueList).(uint)
}
