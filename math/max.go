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

func (m max) Float32(value float32, valueList ...float32) float32 {
	return m.Interface(value, valueList).(float32)
}

func (m max) Float64(value float64, valueList ...float64) float64 {
	return m.Interface(value, valueList).(float64)
}

func (m max) Int(value int, valueList ...int) int {
	return m.Interface(value, valueList).(int)
}

func (m max) Int8(value int8, valueList ...int8) int8 {
	return m.Interface(value, valueList).(int8)
}

func (m max) Int16(value int16, valueList ...int16) int16 {
	return m.Interface(value, valueList).(int16)
}

func (m max) Int32(value int32, valueList ...int32) int32 {
	return m.Interface(value, valueList).(int32)
}

func (m max) Int64(value int64, valueList ...int64) int64 {
	return m.Interface(value, valueList).(int64)
}

func (m max) UInt(value uint, valueList ...uint) uint {
	return m.Interface(value, valueList).(uint)
}

func (m max) UInt8(value uint8, valueList ...uint8) uint8 {
	return m.Interface(value, valueList).(uint8)
}

func (m max) UInt16(value uint16, valueList ...uint16) uint16 {
	return m.Interface(value, valueList).(uint16)
}

func (m max) UInt32(value uint32, valueList ...uint32) uint32 {
	return m.Interface(value, valueList).(uint32)
}

func (m max) UInt64(value uint64, valueList ...uint64) uint64 {
	return m.Interface(value, valueList).(uint64)
}
