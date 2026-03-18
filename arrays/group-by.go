package arrays

func GroupBy[T any, Group comparable](list []T, fn func(value T) Group) map[Group][]T {
	result := make(map[Group][]T)
	for _, value := range list {
		mapKey := fn(value)
		result[mapKey] = append(result[mapKey], value)
	}
	return result
}
