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

func (m min) Float32(value float32, valueList ...float32) float32 {
	return m.Interface(value, valueList).(float32)
}

func (m min) Float64(value float64, valueList ...float64) float64 {
	return m.Interface(value, valueList).(float64)
}

func (m min) Int(value int, valueList ...int) int {
	return m.Interface(value, valueList).(int)
}

func (m min) Int8(value int8, valueList ...int8) int8 {
	return m.Interface(value, valueList).(int8)
}

func (m min) Int16(value int16, valueList ...int16) int16 {
	return m.Interface(value, valueList).(int16)
}

func (m min) Int32(value int32, valueList ...int32) int32 {
	return m.Interface(value, valueList).(int32)
}

func (m min) Int64(value int64, valueList ...int64) int64 {
	return m.Interface(value, valueList).(int64)
}

func (m min) UInt(value uint, valueList ...uint) uint {
	return m.Interface(value, valueList).(uint)
}

func (m min) UInt8(value uint8, valueList ...uint8) uint8 {
	return m.Interface(value, valueList).(uint8)
}

func (m min) UInt16(value uint16, valueList ...uint16) uint16 {
	return m.Interface(value, valueList).(uint16)
}

func (m min) UInt32(value uint32, valueList ...uint32) uint32 {
	return m.Interface(value, valueList).(uint32)
}

func (m min) UInt64(value uint64, valueList ...uint64) uint64 {
	return m.Interface(value, valueList).(uint64)
}
