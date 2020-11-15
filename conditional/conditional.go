package conditional

func Interface(condition bool, trueValue, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}

func Int(condition bool, trueValue, falseValue int) int {
	return Interface(condition, trueValue, falseValue).(int)
}

func String(condition bool, trueValue, falseValue string) string {
	return Interface(condition, trueValue, falseValue).(string)
}
