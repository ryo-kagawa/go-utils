package min

import (
	"math"
	"reflect"
)

func toFloat64(value interface{}) float64 {
	reflectValue := reflect.ValueOf(value)
	switch reflectValue.Kind() {
	case reflect.Float32, reflect.Float64:
		return reflectValue.Float()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(reflectValue.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return float64(reflectValue.Uint())
	default:
		return 0
	}
}

func Interface(value interface{}, valueList interface{}) interface{} {
	result := toFloat64(value)
	valueListReflectValue := reflect.ValueOf(valueList)
	length := valueListReflectValue.Len()
	for i := 0; i < length; i++ {
		result = math.Min(result, toFloat64(valueListReflectValue.Index(i).Interface()))
	}

	return reflect.ValueOf(result).Convert(reflect.ValueOf(value).Type()).Interface()
}

func Float64(value float64, valueList ...float64) float64 {
	return Interface(value, valueList).(float64)
}

func Int(value int, valueList ...int) int {
	return Interface(value, valueList).(int)
}

func UInt(value uint, valueList ...uint) uint {
	return Interface(value, valueList).(uint)
}
