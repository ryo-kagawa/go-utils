package arrays

import (
	"reflect"
	"testing"
)

func Test_groupByImpl_Interface(t *testing.T) {
	zero := 0
	type args struct {
		list interface{}
		fn   interface{}
	}
	tests := []struct {
		name string
		m    groupByImpl
		args args
		want interface{}
	}{
		{
			name: "func([]int(nil)) == map[int][]int{}",
			m:    GroupBy,
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
			m:    GroupBy,
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
			m:    GroupBy,
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
			name: "func([]*int(nil)) == map[int][]*int{}",
			m:    GroupBy,
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
			m:    GroupBy,
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
			if got := tt.m.Interface(tt.args.list, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupByImpl.Interface() = %v, want %v", got, tt.want)
			}
		})
	}
}
