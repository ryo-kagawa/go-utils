package math

import (
	"reflect"
	"testing"
)

func Test_clamp_Interface(t *testing.T) {
	type args struct {
		value interface{}
		min   interface{}
		max   interface{}
	}
	tests := []struct {
		name string
		m    clamp
		args args
		want interface{}
	}{
		{
			name: "func(int(0), int(0), int(0)) == int(0)",
			args: args{
				value: 0,
				min:   0,
				max:   0,
			},
			want: 0,
		},
		{
			name: "func(int(0), int(-1), int(1)) == int(0)",
			args: args{
				value: int(0),
				min:   int(-1),
				max:   int(1),
			},
			want: int(0),
		},
		{
			name: "func(int(0), int(1), int(2)) == int(1)",
			args: args{
				value: int(0),
				min:   int(1),
				max:   int(2),
			},
			want: int(1),
		},
		{
			name: "func(int(0), int(-2), int(-1)) == int(-1)",
			args: args{
				value: int(0),
				min:   int(-2),
				max:   int(-1),
			},
			want: int(-1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Interface(tt.args.value, tt.args.min, tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clamp.Interface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clamp_Float32(t *testing.T) {
	if !reflect.DeepEqual(Clamp.Float32(1, 1, 1), float32(1)) {
		t.Errorf("Clamp.Float32(1, 1, 1) == float32(1)")
	}
}

func Test_clamp_Float64(t *testing.T) {
	if !reflect.DeepEqual(Clamp.Float64(1, 1, 1), float64(1)) {
		t.Errorf("Clamp.Float64(1, 1, 1) == float64(1)")
	}
}

func Test_clamp_Int(t *testing.T) {
	if !reflect.DeepEqual(Clamp.Int(1, 1, 1), int(1)) {
		t.Errorf("Clamp.Int(1, 1, 1) == int(1)")
	}
}

func Test_clamp_Int8(t *testing.T) {
	if !reflect.DeepEqual(Clamp.Int8(1, 1, 1), int8(1)) {
		t.Errorf("Clamp.Int8(1, 1, 1) == int8(1)")
	}
}

func Test_clamp_Int16(t *testing.T) {
	if !reflect.DeepEqual(Clamp.Int16(1, 1, 1), int16(1)) {
		t.Errorf("Clamp.Int16(1, 1, 1) == int16(1)")
	}
}

func Test_clamp_Int32(t *testing.T) {
	if !reflect.DeepEqual(Clamp.Int32(1, 1, 1), int32(1)) {
		t.Errorf("Clamp.Int32(1, 1, 1) == int32(1)")
	}
}

func Test_clamp_Int64(t *testing.T) {
	if !reflect.DeepEqual(Clamp.Int64(1, 1, 1), int64(1)) {
		t.Errorf("Clamp.Int64(1, 1, 1) == int64(1)")
	}
}

func Test_clamp_UInt(t *testing.T) {
	if !reflect.DeepEqual(Clamp.UInt(1, 1, 1), uint(1)) {
		t.Errorf("Clamp.UInt(1, 1, 1) == uint(1)")
	}
}

func Test_clamp_UInt8(t *testing.T) {
	if !reflect.DeepEqual(Clamp.UInt8(1, 1, 1), uint8(1)) {
		t.Errorf("Clamp.UInt8(1, 1, 1) == uint8(1)")
	}
}

func Test_clamp_UInt16(t *testing.T) {
	if !reflect.DeepEqual(Clamp.UInt16(1, 1, 1), uint16(1)) {
		t.Errorf("Clamp.UInt16(1, 1, 1) == uint16(1)")
	}
}

func Test_clamp_UInt32(t *testing.T) {
	if !reflect.DeepEqual(Clamp.UInt32(1, 1, 1), uint32(1)) {
		t.Errorf("Clamp.UInt32(1, 1, 1) == uint32(1)")
	}
}

func Test_clamp_UInt64(t *testing.T) {
	if !reflect.DeepEqual(Clamp.UInt64(1, 1, 1), uint64(1)) {
		t.Errorf("Clamp.UInt64(1, 1, 1) == uint64(1)")
	}
}
