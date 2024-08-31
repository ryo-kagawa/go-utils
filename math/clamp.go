package math

func Clamp[T Number](value, min, max T) T {
	if value <= min {
		return min
	}
	if max <= value {
		return max
	}
	return value
}
