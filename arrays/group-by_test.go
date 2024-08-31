package arrays

import (
	"reflect"
	"testing"
)

func TestGroupBy_int_int(t *testing.T) {
	type args struct {
		list []int
		fn   func(value int) int
	}
	tests := []struct {
		name string
		args args
		want map[int][]int
	}{
		{
			name: "func([]int(nil)) == map[int][]int{}",
			args: args{
				list: []int(nil),
				fn: func(v int) int {
					return v
				},
			},
			want: map[int][]int{},
		},
		{
			name: "func([]int{0}) == map[int][]int{0: {0}}",
			args: args{
				list: []int{0},
				fn: func(v int) int {
					return v
				},
			},
			want: map[int][]int{
				0: {0},
			},
		},
		{
			name: "func([]int{0, 1, 0}) == map[int][]int{0: {0, 0}, 1: {1}}",
			args: args{
				list: []int{0, 1, 0},
				fn: func(v int) int {
					return v
				},
			},
			want: map[int][]int{
				0: {0, 0},
				1: {1},
			},
		},
		{
			name: "func([]int{0, 1, 0}) == map[int][]int{0: {0, 0}, 1: {1}}",
			args: args{
				list: []int{0, 1, 0},
				fn: func(v int) int {
					return v
				},
			},
			want: map[int][]int{
				0: {0, 0},
				1: {1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupBy(tt.args.list, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupBy_intpointer_int(t *testing.T) {
	zero := 0
	type args struct {
		list []*int
		fn   func(value *int) int
	}
	tests := []struct {
		name string
		args args
		want map[int][]*int
	}{
		{
			name: "func([]*int(nil)) == map[int][]*int{}",
			args: args{
				list: []*int(nil),
				fn: func(v *int) int {
					return *v
				},
			},
			want: map[int][]*int{},
		},
		{
			name: "func([]*int{0}) == map[int][]*int{0: {0}}",
			args: args{
				list: []*int{&zero},
				fn: func(v *int) int {
					return *v
				},
			},
			want: map[int][]*int{
				0: {&zero},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupBy(tt.args.list, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
