package binary

import "bytes"

func CounterBigEndian(value []byte) []byte {
	value = bytes.Clone(value)
	for i := len(value) - 1; 0 <= i; i-- {
		value[i]++
		if value[i] != 0 {
			return value
		}
	}
	return append([]byte{0x01}, value...)
}

func CounterLittleEndian(value []byte) []byte {
	value = bytes.Clone(value)
	for i := range value {
		value[i]++
		if value[i] != 0 {
			return value
		}
	}
	return append(value, 0x01)
}
