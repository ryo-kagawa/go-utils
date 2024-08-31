package conditional

func Func[T any](condition bool, trueFunc, falseFunc func() T) T {
	if condition {
		return trueFunc()
	}
	return falseFunc()
}
