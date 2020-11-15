package conditional

func InterfaceFunc(condition bool, trueFunc, falseFunc func() interface{}) interface{} {
	if condition {
		return trueFunc()
	}
	return falseFunc()
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
