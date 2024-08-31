package math

import (
	"testing"
)

func TestInRange_int(t *testing.T) {
	type args struct {
		value int
		min   int
		max   int
	}
	tests := []struct {
		name string
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
			if got := InRange(tt.args.value, tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("InRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
