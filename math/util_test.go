package math

import "testing"

func TestToFloat64(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "func(float32(1)) == float64(1)",
			args: args{
				value: float32(1),
			},
			want: float64(1),
		},
		{
			name: "func(float64(1)) == float64(1)",
			args: args{
				value: float64(1),
			},
			want: float64(1),
		},
		{
			name: "func(int(1)) == float64(1)",
			args: args{
				value: int(1),
			},
			want: float64(1),
		},
		{
			name: "func(int8(1)) == float64(1)",
			args: args{
				value: int8(1),
			},
			want: float64(1),
		},
		{
			name: "func(int16(1)) == float64(1)",
			args: args{
				value: int16(1),
			},
			want: float64(1),
		},
		{
			name: "func(int32(1)) == float64(1)",
			args: args{
				value: int32(1),
			},
			want: float64(1),
		},
		{
			name: "func(int64(1)) == float64(1)",
			args: args{
				value: int64(1),
			},
			want: float64(1),
		},
		{
			name: "func(uint(1)) == float64(1)",
			args: args{
				value: uint(1),
			},
			want: float64(1),
		},
		{
			name: "func(uint8(1)) == float64(1)",
			args: args{
				value: uint8(1),
			},
			want: float64(1),
		},
		{
			name: "func(uint16(1)) == float64(1)",
			args: args{
				value: uint16(1),
			},
			want: float64(1),
		},
		{
			name: "func(uint32(1)) == float64(1)",
			args: args{
				value: uint32(1),
			},
			want: float64(1),
		},
		{
			name: "func(uint64(1)) == float64(1)",
			args: args{
				value: uint64(1),
			},
			want: float64(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFloat64(tt.args.value); got != tt.want {
				t.Errorf("ToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}
