package arrays_test

import (
	"testing"

	"github.com/ryo-kagawa/go-utils/arrays"
)

func TestEvery(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    []bool
		f    func(bool) bool
		want bool
	}{
		{
			name: "func([]bool(nil), func(value bool) { return value }) == true",
			s:    []bool(nil),
			f:    func(value bool) bool { return value },
			want: true,
		},
		{
			name: "func([]bool{}, func(value bool) { return value }) == true",
			s:    []bool{},
			f:    func(value bool) bool { return value },
			want: true,
		},
		{
			name: "func([]bool{true}, func(value bool) { return value }) == true",
			s:    []bool{true},
			f:    func(value bool) bool { return value },
			want: true,
		},
		{
			name: "func([]bool{false}, func(value bool) { return value }) == false",
			s:    []bool{false},
			f:    func(value bool) bool { return value },
			want: false,
		},
		{
			name: "func([]bool{true, true}, func(value bool) { return value }) == true",
			s:    []bool{true, true},
			f:    func(value bool) bool { return value },
			want: true,
		},
		{
			name: "func([]bool{true, false}, func(value bool) { return value }) == false",
			s:    []bool{true, false},
			f:    func(value bool) bool { return value },
			want: false,
		},
		{
			name: "func([]bool{false, true}, func(value bool) { return value }) == false",
			s:    []bool{false, true},
			f:    func(value bool) bool { return value },
			want: false,
		},
		{
			name: "func([]bool{false, false}, func(value bool) { return value }) == false",
			s:    []bool{false, false},
			f:    func(value bool) bool { return value },
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrays.Every(tt.s, tt.f); got != tt.want {
				t.Errorf("Every() = %v, want %v", got, tt.want)
			}
		})
	}
}
