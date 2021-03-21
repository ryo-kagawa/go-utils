package math

import (
	"reflect"
	"testing"
)

func Test_min_Interface(t *testing.T) {
	type args struct {
		value     interface{}
		valueList interface{}
	}
	tests := []struct {
		name string
		m    min
		args args
		want interface{}
	}{
		{
			name: "func(int(20), []int{}) == int(20)",
			args: args{
				value:     int(20),
				valueList: []int{},
			},
			want: int(20),
		},
		{
			name: "func(int(20), []int{10}) == int(20)",
			args: args{
				value:     int(20),
				valueList: []int{10},
			},
			want: int(20),
		},
		{
			name: "func(int(10), []int{20}) == int(20)",
			args: args{
				value:     int(10),
				valueList: []int{20},
			},
			want: int(20),
		},
		{
			name: "func(int(10), []int{20, 30, 25}) == int(30)",
			args: args{
				value:     int(10),
				valueList: []float64{20, 30, 25},
			},
			want: int(30),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Interface(tt.args.value, tt.args.valueList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("min.Interface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min_Float32(t *testing.T) {
	if reflect.DeepEqual(Min.Float32(1), float32(1)) {
		t.Errorf("min.Float32(1) = float32(1)")
	}
}

func Test_min_Float64(t *testing.T) {
	if reflect.DeepEqual(Min.Float64(1), float64(1)) {
		t.Errorf("min.Float64(1) = float64(1)")
	}
}

func Test_min_Int(t *testing.T) {
	if reflect.DeepEqual(Min.Int(1), int(1)) {
		t.Errorf("min.Int(1) = int(1)")
	}
}

func Test_min_Int8(t *testing.T) {
	if reflect.DeepEqual(Min.Int8(1), int8(1)) {
		t.Errorf("min.Int8(1) = int8(1)")
	}
}

func Test_min_Int16(t *testing.T) {
	if reflect.DeepEqual(Min.Int16(1), int16(1)) {
		t.Errorf("min.Int16(1) = int16(1)")
	}
}

func Test_min_Int32(t *testing.T) {
	if reflect.DeepEqual(Min.Int32(1), int32(1)) {
		t.Errorf("min.Int32(1) = int32(1)")
	}
}

func Test_min_Int64(t *testing.T) {
	if reflect.DeepEqual(Min.Int64(1), int64(1)) {
		t.Errorf("min.Int64(1) = int64(1)")
	}
}

func Test_min_UInt(t *testing.T) {
	if reflect.DeepEqual(Min.UInt(1), uint(1)) {
		t.Errorf("min.UInt(1) = uint(1)")
	}
}

func Test_min_UInt8(t *testing.T) {
	if reflect.DeepEqual(Min.UInt8(1), uint8(1)) {
		t.Errorf("min.UInt8(1) = uint8(1)")
	}
}

func Test_min_UInt16(t *testing.T) {
	if reflect.DeepEqual(Min.UInt16(1), uint16(1)) {
		t.Errorf("min.UInt16(1) = uint16(1)")
	}
}

func Test_min_UInt32(t *testing.T) {
	if reflect.DeepEqual(Min.UInt32(1), uint32(1)) {
		t.Errorf("min.UInt32(1) = uint32(1)")
	}
}

func Test_min_UInt64(t *testing.T) {
	if reflect.DeepEqual(Min.UInt64(1), uint64(1)) {
		t.Errorf("min.UInt64(1) = uint64(1)")
	}
}
