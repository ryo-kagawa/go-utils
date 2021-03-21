package math

import "reflect"

func ToFloat64(value interface{}) float64 {
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
