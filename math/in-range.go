package math

type inRange struct{}

var InRange = inRange{}

func (m inRange) Interface(value interface{}, min interface{}, max interface{}) bool {
	return ToFloat64(min) <= ToFloat64(value) && ToFloat64(value) <= ToFloat64(max)
}

func (m inRange) Float32(value, min, max float32) bool {
	return m.Interface(value, min, max)
}

func (m inRange) Float64(value, min, max float64) bool {
	return m.Interface(value, min, max)
}

func (m inRange) Int(value, min, max int) bool {
	return m.Interface(value, min, max)
}

func (m inRange) Int8(value, min, max int8) bool {
	return m.Interface(value, min, max)
}

func (m inRange) Int16(value, min, max int16) bool {
	return m.Interface(value, min, max)
}

func (m inRange) Int32(value, min, max int32) bool {
	return m.Interface(value, min, max)
}

func (m inRange) Int64(value, min, max int64) bool {
	return m.Interface(value, min, max)
}

func (m inRange) UInt(value, min, max uint) bool {
	return m.Interface(value, min, max)
}

func (m inRange) UInt8(value, min, max uint8) bool {
	return m.Interface(value, min, max)
}

func (m inRange) UInt16(value, min, max uint16) bool {
	return m.Interface(value, min, max)
}

func (m inRange) UInt32(value, min, max uint32) bool {
	return m.Interface(value, min, max)
}

func (m inRange) UInt64(value, min, max uint64) bool {
	return m.Interface(value, min, max)
}
