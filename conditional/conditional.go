package conditional

func Interface(condition bool, trueValue, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}

func Bool(condition bool, trueValue, falseValue bool) bool {
	return Interface(condition, trueValue, falseValue).(bool)
}

func Float32(condition bool, trueValue, falseValue float32) float32 {
	return Interface(condition, trueValue, falseValue).(float32)
}

func Float64(condition bool, trueValue, falseValue float64) float64 {
	return Interface(condition, trueValue, falseValue).(float64)
}

func Int(condition bool, trueValue, falseValue int) int {
	return Interface(condition, trueValue, falseValue).(int)
}

func Int8(condition bool, trueValue, falseValue int8) int8 {
	return Interface(condition, trueValue, falseValue).(int8)
}

func Int16(condition bool, trueValue, falseValue int16) int16 {
	return Interface(condition, trueValue, falseValue).(int16)
}

func Int32(condition bool, trueValue, falseValue int32) int32 {
	return Interface(condition, trueValue, falseValue).(int32)
}

func Int64(condition bool, trueValue, falseValue int64) int64 {
	return Interface(condition, trueValue, falseValue).(int64)
}

func String(condition bool, trueValue, falseValue string) string {
	return Interface(condition, trueValue, falseValue).(string)
}

func UInt(condition bool, trueValue, falseValue uint) uint {
	return Interface(condition, trueValue, falseValue).(uint)
}

func UInt8(condition bool, trueValue, falseValue uint8) uint8 {
	return Interface(condition, trueValue, falseValue).(uint8)
}

func UInt16(condition bool, trueValue, falseValue uint16) uint16 {
	return Interface(condition, trueValue, falseValue).(uint16)
}

func UInt32(condition bool, trueValue, falseValue uint32) uint32 {
	return Interface(condition, trueValue, falseValue).(uint32)
}

func UInt64(condition bool, trueValue, falseValue uint64) uint64 {
	return Interface(condition, trueValue, falseValue).(uint64)
}
