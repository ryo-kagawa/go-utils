package math

import (
	"math"
	"reflect"
	"testing"
)

func Test_min_Float32(t *testing.T) {
	if !reflect.DeepEqual(Min.Float32(0, -1), float32(-1)) {
		t.Errorf("max.Float32(0, -1) = float32(-1)")
	}
}

func Test_min_Float64(t *testing.T) {
	if !reflect.DeepEqual(Min.Float64(0, -1), float64(-1)) {
		t.Errorf("min.Float64(0, -1) = float64(-1)")
	}
}

func Test_min_Int(t *testing.T) {
	if !reflect.DeepEqual(Min.Int(0, -1), int(-1)) {
		t.Errorf("min.Int(0, -1) = int(-1)")
	}
}

func Test_min_Int8(t *testing.T) {
	if !reflect.DeepEqual(Min.Int8(0, math.MinInt8), int8(math.MinInt8)) {
		t.Errorf("min.Int8(0, math.MinInt8) = int8(math.MinInt8)")
	}
}

func Test_min_Int16(t *testing.T) {
	if !reflect.DeepEqual(Min.Int16(0, math.MinInt16), int16(math.MinInt16)) {
		t.Errorf("min.Int16(0, math.MinInt16) = int16(math.MinInt16)")
	}
}

func Test_min_Int32(t *testing.T) {
	if !reflect.DeepEqual(Min.Int32(0, math.MinInt32), int32(math.MinInt32)) {
		t.Errorf("min.Int32(0, math.MinInt32) = int32(math.MinInt32)")
	}
}

func Test_min_Int64(t *testing.T) {
	if !reflect.DeepEqual(Min.Int64(0, math.MinInt64), int64(math.MinInt64)) {
		t.Errorf("min.Int64(0, math.MinInt64) = int64(math.MinInt64)")
	}
}

func Test_min_UInt(t *testing.T) {
	if !reflect.DeepEqual(Min.UInt(0, 1), uint(0)) {
		t.Errorf("min.UInt(0, 1) = uint(0)")
	}
}

func Test_min_UInt8(t *testing.T) {
	if !reflect.DeepEqual(Min.UInt8(0, 1), uint8(0)) {
		t.Errorf("min.UInt8(0, 1) = uint8(1)")
	}
}

func Test_min_UInt16(t *testing.T) {
	if !reflect.DeepEqual(Min.UInt16(0, 1), uint16(0)) {
		t.Errorf("min.UInt16(0, 1) = uint16(1)")
	}
}

func Test_min_UInt32(t *testing.T) {
	if !reflect.DeepEqual(Min.UInt32(0, 1), uint32(0)) {
		t.Errorf("min.UInt32(0, 1) = uint32(1)")
	}
}

func Test_min_UInt64(t *testing.T) {
	if !reflect.DeepEqual(Min.UInt64(0, 1), uint64(0)) {
		t.Errorf("min.UInt64(0, 1) = uint64(1)")
	}
}
