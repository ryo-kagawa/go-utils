package math

func InRange[T Number](value, min, max T) bool {
	return min <= value && value <= max
}
