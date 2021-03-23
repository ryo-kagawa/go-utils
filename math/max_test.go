package math

import (
	"reflect"
	"testing"
)

func Test_max_Interface(t *testing.T) {
	type args struct {
		value     interface{}
		valueList interface{}
	}
	tests := []struct {
		name string
		m    max
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
				t.Errorf("max.Interface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_max_Float32(t *testing.T) {
	if !reflect.DeepEqual(Max.Float32(1), float32(1)) {
		t.Errorf("max.Float32(1) = float32(1)")
	}
}

func Test_max_Float64(t *testing.T) {
	if !reflect.DeepEqual(Max.Float64(1), float64(1)) {
		t.Errorf("max.Float64(1) = float64(1)")
	}
}

func Test_max_Int(t *testing.T) {
	if !reflect.DeepEqual(Max.Int(1), int(1)) {
		t.Errorf("max.Int(1) = int(1)")
	}
}

func Test_max_Int8(t *testing.T) {
	if !reflect.DeepEqual(Max.Int8(1), int8(1)) {
		t.Errorf("max.Int8(1) = int8(1)")
	}
}

func Test_max_Int16(t *testing.T) {
	if !reflect.DeepEqual(Max.Int16(1), int16(1)) {
		t.Errorf("max.Int16(1) = int16(1)")
	}
}

func Test_max_Int32(t *testing.T) {
	if !reflect.DeepEqual(Max.Int32(1), int32(1)) {
		t.Errorf("max.Int32(1) = int32(1)")
	}
}

func Test_max_Int64(t *testing.T) {
	if !reflect.DeepEqual(Max.Int64(1), int64(1)) {
		t.Errorf("max.Int64(1) = int64(1)")
	}
}

func Test_max_UInt(t *testing.T) {
	if !reflect.DeepEqual(Max.UInt(1), uint(1)) {
		t.Errorf("max.UInt(1) = uint(1)")
	}
}

func Test_max_UInt8(t *testing.T) {
	if !reflect.DeepEqual(Max.UInt8(1), uint8(1)) {
		t.Errorf("max.UInt8(1) = uint8(1)")
	}
}

func Test_max_UInt16(t *testing.T) {
	if !reflect.DeepEqual(Max.UInt16(1), uint16(1)) {
		t.Errorf("max.UInt16(1) = uint16(1)")
	}
}

func Test_max_UInt32(t *testing.T) {
	if !reflect.DeepEqual(Max.UInt32(1), uint32(1)) {
		t.Errorf("max.UInt32(1) = uint32(1)")
	}
}

func Test_max_UInt64(t *testing.T) {
	if !reflect.DeepEqual(Max.UInt64(1), uint64(1)) {
		t.Errorf("max.UInt64(1) = uint64(1)")
	}
}
