package arrays

func Reduce[T any, ResultType any](list []T, fn func(accumulator ResultType, value T, currentIndex int, array []T) ResultType, initialValue ResultType) ResultType {
	accumulator := initialValue
	for index, value := range list {
		accumulator = fn(accumulator, value, index, list)
	}
	return accumulator
}
