package math

import "testing"

func Test_inRange_Interface(t *testing.T) {
	type args struct {
		value interface{}
		min   interface{}
		max   interface{}
	}
	tests := []struct {
		name string
		m    inRange
		args args
		want bool
	}{
		{
			name: "func(0, 0, 0) == true",
			args: args{
				value: 0,
				min:   0,
				max:   0,
			},
			want: true,
		},
		{
			name: "func(1, 0, 0) == false",
			args: args{
				value: 1,
				min:   0,
				max:   0,
			},
			want: false,
		},
		{
			name: "func(1, 0, 1) == true",
			args: args{
				value: 1,
				min:   0,
				max:   1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Interface(tt.args.value, tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("inRange.Interface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inRange_Float32(t *testing.T) {
	if !InRange.Float32(float32(1), float32(1), float32(1)) {
		t.Errorf("InRange.Float32(float32(1), float32(1), float32(1)")
	}
}

func Test_inRange_Float64(t *testing.T) {
	if !InRange.Float64(float64(1), float64(1), float64(1)) {
		t.Errorf("InRange.Float64(float64(1), float64(1), float64(1)")
	}
}

func Test_inRange_Int(t *testing.T) {
	if !InRange.Int(int(1), int(1), int(1)) {
		t.Errorf("InRange.Int(int(1), int(1), int(1)")
	}
}

func Test_inRange_Int8(t *testing.T) {
	if !InRange.Int8(int8(1), int8(1), int8(1)) {
		t.Errorf("InRange.Int8(int8(1), int8(1), int8(1)")
	}
}

func Test_inRange_Int16(t *testing.T) {
	if !InRange.Int16(int16(1), int16(1), int16(1)) {
		t.Errorf("InRange.Int16(int16(1), int16(1), int16(1)")
	}
}

func Test_inRange_Int32(t *testing.T) {
	if !InRange.Int32(int32(1), int32(1), int32(1)) {
		t.Errorf("InRange.Int32(int32(1), int32(1), int32(1)")
	}
}

func Test_inRange_Int64(t *testing.T) {
	if !InRange.Int64(int64(1), int64(1), int64(1)) {
		t.Errorf("InRange.Int64(int64(1), int64(1), int64(1)")
	}
}

func Test_inRange_UInt(t *testing.T) {
	if !InRange.UInt(uint(1), uint(1), uint(1)) {
		t.Errorf("InRange.UInt(uint(1), uint(1), uint(1)")
	}
}

func Test_inRange_UInt8(t *testing.T) {
	if !InRange.UInt8(uint8(1), uint8(1), uint8(1)) {
		t.Errorf("InRange.UInt8(uint8(1), uint8(1), uint8(1)")
	}
}

func Test_inRange_UInt16(t *testing.T) {
	if !InRange.UInt16(uint16(1), uint16(1), uint16(1)) {
		t.Errorf("InRange.UInt16(uint16(1), uint16(1), uint16(1)")
	}
}

func Test_inRange_UInt32(t *testing.T) {
	if !InRange.UInt32(uint32(1), uint32(1), uint32(1)) {
		t.Errorf("InRange.UInt32(uint32(1), uint32(1), uint32(1)")
	}
}

func Test_inRange_UInt64(t *testing.T) {
	if !InRange.UInt64(uint64(1), uint64(1), uint64(1)) {
		t.Errorf("InRange.UInt64(uint64(1), uint64(1), uint64(1)")
	}
}
