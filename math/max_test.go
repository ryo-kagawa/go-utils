package math

import (
	"math"
	"reflect"
	"testing"
)

func Test_max_Float32(t *testing.T) {
	if !reflect.DeepEqual(Max.Float32(0, math.MaxFloat32), float32(math.MaxFloat32)) {
		t.Errorf("max.Float32(0, math.MaxFloat32) = float32(math.MaxFloat32)")
	}
}

func Test_max_Float64(t *testing.T) {
	if !reflect.DeepEqual(Max.Float64(0, math.MaxFloat64), float64(math.MaxFloat64)) {
		t.Errorf("max.Float64(0, math.MaxFloat64) = float64(math.MaxFloat64)")
	}
}

func Test_max_Int(t *testing.T) {
	if !reflect.DeepEqual(Max.Int(0, 1), int(1)) {
		t.Errorf("max.Int(0, 1) = int(1)")
	}
}

func Test_max_Int8(t *testing.T) {
	if !reflect.DeepEqual(Max.Int8(0, math.MaxInt8), int8(math.MaxInt8)) {
		t.Errorf("max.Int8(0, math.MaxInt8) = int8(math.MaxInt8)")
	}
}

func Test_max_Int16(t *testing.T) {
	if !reflect.DeepEqual(Max.Int16(0, math.MaxInt16), int16(math.MaxInt16)) {
		t.Errorf("max.Int16(0, math.MaxInt16) = int16(math.MaxInt16)")
	}
}

func Test_max_Int32(t *testing.T) {
	if !reflect.DeepEqual(Max.Int32(0, math.MaxInt32), int32(math.MaxInt32)) {
		t.Errorf("max.Int32(0, math.MaxInt32) = int32(math.MaxInt32)")
	}
}

func Test_max_Int64(t *testing.T) {
	if !reflect.DeepEqual(Max.Int64(0, math.MaxInt64), int64(math.MaxInt64)) {
		t.Errorf("max.Int64(0, math.MaxInt64) = int64(math.MaxInt64)")
	}
}

func Test_max_UInt(t *testing.T) {
	if !reflect.DeepEqual(Max.UInt(0, 1), uint(1)) {
		t.Errorf("max.UInt(0, 1) = uint(1)")
	}
}

func Test_max_UInt8(t *testing.T) {
	if !reflect.DeepEqual(Max.UInt8(0, math.MaxUint8), uint8(math.MaxUint8)) {
		t.Errorf("max.UInt8(0, math.MaxUint8) = uint8(math.MaxUint8)")
	}
}

func Test_max_UInt16(t *testing.T) {
	if !reflect.DeepEqual(Max.UInt16(0, math.MaxUint16), uint16(math.MaxUint16)) {
		t.Errorf("max.UInt16(0, math.MaxUint16) = uint16(math.MaxUint16)")
	}
}

func Test_max_UInt32(t *testing.T) {
	if !reflect.DeepEqual(Max.UInt32(0, math.MaxUint32), uint32(math.MaxUint32)) {
		t.Errorf("max.UInt32(0, math.MaxUint32) = uint32(math.MaxUint32)")
	}
}

func Test_max_UInt64(t *testing.T) {
	if !reflect.DeepEqual(Max.UInt64(0, math.MaxUint64), uint64(math.MaxUint64)) {
		t.Errorf("max.UInt64(0, math.MaxUint64) = uint64(math.MaxUint64)")
	}
}
