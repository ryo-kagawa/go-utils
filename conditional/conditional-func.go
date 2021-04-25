package conditional

func InterfaceFunc(condition bool, trueFunc, falseFunc func() interface{}) interface{} {
	if condition {
		return trueFunc()
	}
	return falseFunc()
}

func BoolFunc(condition bool, trueFunc, falseFunc func() bool) bool {
	return InterfaceFunc(
		condition,
		func() interface{} {
			return trueFunc()
		},
		func() interface{} {
			return falseFunc()
		},
	).(bool)
}

func Float32Func(condition bool, trueFunc, falseFunc func() float32) float32 {
	return InterfaceFunc(
		condition,
		func() interface{} {
			return trueFunc()
		},
		func() interface{} {
			return falseFunc()
		},
	).(float32)
}

func Float64Func(condition bool, trueFunc, falseFunc func() float64) float64 {
	return InterfaceFunc(
		condition,
		func() interface{} {
			return trueFunc()
		},
		func() interface{} {
			return falseFunc()
		},
	).(float64)
}

func IntFunc(condition bool, trueFunc, falseFunc func() int) int {
	return InterfaceFunc(
		condition,
		func() interface{} {
			return trueFunc()
		},
		func() interface{} {
			return falseFunc()
		},
	).(int)
}

func Int8Func(condition bool, trueFunc, falseFunc func() int8) int8 {
	return InterfaceFunc(
		condition,
		func() interface{} {
			return trueFunc()
		},
		func() interface{} {
			return falseFunc()
		},
	).(int8)
}

func Int16Func(condition bool, trueFunc, falseFunc func() int16) int16 {
	return InterfaceFunc(
		condition,
		func() interface{} {
			return trueFunc()
		},
		func() interface{} {
			return falseFunc()
		},
	).(int16)
}

func Int32Func(condition bool, trueFunc, falseFunc func() int32) int32 {
	return InterfaceFunc(
		condition,
		func() interface{} {
			return trueFunc()
		},
		func() interface{} {
			return falseFunc()
		},
	).(int32)
}

func Int64Func(condition bool, trueFunc, falseFunc func() int64) int64 {
	return InterfaceFunc(
		condition,
		func() interface{} {
			return trueFunc()
		},
		func() interface{} {
			return falseFunc()
		},
	).(int64)
}

func StringFunc(condition bool, trueFunc, falseFunc func() string) string {
	return InterfaceFunc(
		condition,
		func() interface{} {
			return trueFunc()
		},
		func() interface{} {
			return falseFunc()
		},
	).(string)
}

func UIntFunc(condition bool, trueFunc, falseFunc func() uint) uint {
	return InterfaceFunc(
		condition,
		func() interface{} {
			return trueFunc()
		},
		func() interface{} {
			return falseFunc()
		},
	).(uint)
}

func UInt8Func(condition bool, trueFunc, falseFunc func() uint8) uint8 {
	return InterfaceFunc(
		condition,
		func() interface{} {
			return trueFunc()
		},
		func() interface{} {
			return falseFunc()
		},
	).(uint8)
}

func UInt16Func(condition bool, trueFunc, falseFunc func() uint16) uint16 {
	return InterfaceFunc(
		condition,
		func() interface{} {
			return trueFunc()
		},
		func() interface{} {
			return falseFunc()
		},
	).(uint16)
}

func UInt32Func(condition bool, trueFunc, falseFunc func() uint32) uint32 {
	return InterfaceFunc(
		condition,
		func() interface{} {
			return trueFunc()
		},
		func() interface{} {
			return falseFunc()
		},
	).(uint32)
}

func UInt64Func(condition bool, trueFunc, falseFunc func() uint64) uint64 {
	return InterfaceFunc(
		condition,
		func() interface{} {
			return trueFunc()
		},
		func() interface{} {
			return falseFunc()
		},
	).(uint64)
}
