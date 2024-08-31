package math

import (
	"reflect"
	"testing"
)

func TestClamp_int(t *testing.T) {
	type args struct {
		value int
		min   int
		max   int
	}
	tests := []struct {
		name string
		args args
		want int
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
				value: 0,
				min:   -1,
				max:   1,
			},
			want: 0,
		},
		{
			name: "func(int(0), int(1), int(2)) == int(1)",
			args: args{
				value: 0,
				min:   1,
				max:   2,
			},
			want: 1,
		},
		{
			name: "func(int(0), int(-2), int(-1)) == int(-1)",
			args: args{
				value: 0,
				min:   -2,
				max:   -1,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clamp(tt.args.value, tt.args.min, tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
