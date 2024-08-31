package arrays

func GroupBy[T any, Group comparable](list []T, fn func(value T) Group) map[Group][]T {
	result := make(map[Group][]T, 0)
	for _, value := range list {
		mapKey := fn(value)
		mapValue, ok := result[mapKey]
		if !ok {
			result[mapKey] = []T{}
		}
		result[mapKey] = append(mapValue, value)
	}
	return result
}
