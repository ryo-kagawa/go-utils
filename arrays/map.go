package arrays

func Map[T any, ResponseType any](list []T, fn func(value T) ResponseType) []ResponseType {
	result := make([]ResponseType, 0, len(list))
	for _, value := range list {
		result = append(result, fn(value))
	}
	return result
}
