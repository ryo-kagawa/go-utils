package math

import (
	"reflect"
)

type clamp struct{}

var Clamp = clamp{}

func (m clamp) Interface(value interface{}, min interface{}, max interface{}) interface{} {
	return reflect.ValueOf(Min.Float64(Max.Float64(ToFloat64(value), ToFloat64(min)), ToFloat64(max))).Convert(reflect.ValueOf(value).Type()).Interface()
}

func (m clamp) Float32(value, min, max float32) float32 {
	return m.Interface(value, min, max).(float32)
}

func (m clamp) Float64(value, min, max float64) float64 {
	return m.Interface(value, min, max).(float64)
}

func (m clamp) Int(value, min, max int) int {
	return m.Interface(value, min, max).(int)
}

func (m clamp) Int8(value, min, max int8) int8 {
	return m.Interface(value, min, max).(int8)
}

func (m clamp) Int16(value, min, max int16) int16 {
	return m.Interface(value, min, max).(int16)
}

func (m clamp) Int32(value, min, max int32) int32 {
	return m.Interface(value, min, max).(int32)
}

func (m clamp) Int64(value, min, max int64) int64 {
	return m.Interface(value, min, max).(int64)
}

func (m clamp) UInt(value, min, max uint) uint {
	return m.Interface(value, min, max).(uint)
}

func (m clamp) UInt8(value, min, max uint8) uint8 {
	return m.Interface(value, min, max).(uint8)
}

func (m clamp) UInt16(value, min, max uint16) uint16 {
	return m.Interface(value, min, max).(uint16)
}

func (m clamp) UInt32(value, min, max uint32) uint32 {
	return m.Interface(value, min, max).(uint32)
}

func (m clamp) UInt64(value, min, max uint64) uint64 {
	return m.Interface(value, min, max).(uint64)
}
