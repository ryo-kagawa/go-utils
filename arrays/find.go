package arrays

func Find[T any](list []T, fn func(value T) bool) (T, bool) {
	for _, value := range list {
		if fn(value) {
			return value, true
		}
	}
	var ret T
	return ret, false
}
