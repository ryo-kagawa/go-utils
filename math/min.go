package math

type min struct{}

var Min = min{}

func (m min) Float32(value float32, valueList ...float32) float32 {
	result := value
	for _, x := range valueList {
		if result > x {
			result = x
		}
	}
	return result
}

func (m min) Float64(value float64, valueList ...float64) float64 {
	result := value
	for _, x := range valueList {
		if result > x {
			result = x
		}
	}
	return result
}

func (m min) Int(value int, valueList ...int) int {
	result := value
	for _, x := range valueList {
		if result > x {
			result = x
		}
	}
	return result
}

func (m min) Int8(value int8, valueList ...int8) int8 {
	result := value
	for _, x := range valueList {
		if result > x {
			result = x
		}
	}
	return result
}

func (m min) Int16(value int16, valueList ...int16) int16 {
	result := value
	for _, x := range valueList {
		if result > x {
			result = x
		}
	}
	return result
}

func (m min) Int32(value int32, valueList ...int32) int32 {
	result := value
	for _, x := range valueList {
		if result > x {
			result = x
		}
	}
	return result
}

func (m min) Int64(value int64, valueList ...int64) int64 {
	result := value
	for _, x := range valueList {
		if result > x {
			result = x
		}
	}
	return result
}

func (m min) UInt(value uint, valueList ...uint) uint {
	result := value
	for _, x := range valueList {
		if result > x {
			result = x
		}
	}
	return result
}

func (m min) UInt8(value uint8, valueList ...uint8) uint8 {
	result := value
	for _, x := range valueList {
		if result > x {
			result = x
		}
	}
	return result
}

func (m min) UInt16(value uint16, valueList ...uint16) uint16 {
	result := value
	for _, x := range valueList {
		if result > x {
			result = x
		}
	}
	return result
}

func (m min) UInt32(value uint32, valueList ...uint32) uint32 {
	result := value
	for _, x := range valueList {
		if result > x {
			result = x
		}
	}
	return result
}

func (m min) UInt64(value uint64, valueList ...uint64) uint64 {
	result := value
	for _, x := range valueList {
		if result > x {
			result = x
		}
	}
	return result
}
