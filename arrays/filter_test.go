package arrays

import (
	"reflect"
	"testing"
)

func Test_filter_Interface(t *testing.T) {
	zero := 0
	one := 1

	type args struct {
		list interface{}
		fn   interface{}
	}
	tests := []struct {
		name string
		f    filter
		args args
		want interface{}
	}{
		{
			name: "func([]int(nil), true) == []int{}",
			args: args{
				list: []int(nil),
				fn: func(v int) bool {
					return true
				},
			},
			want: []int{},
		},
		{
			name: "func([]int{}, true) == []int{}",
			args: args{
				list: []int{},
				fn: func(v int) bool {
					return true
				},
			},
			want: []int{},
		},
		{
			name: "func([]int{0}, v == 0) == []int{0}",
			args: args{
				list: []int{0},
				fn: func(v int) bool {
					return v == 0
				},
			},
			want: []int{0},
		},
		{
			name: "func([]int{1}, v == 0) == []int{}",
			args: args{
				list: []int{1},
				fn: func(v int) bool {
					return v == 0
				},
			},
			want: []int{},
		},
		{
			name: "func([]int{0, 1, 2}, v != 1) == []int{0, 2}",
			args: args{
				list: []int{0, 1, 2},
				fn: func(v int) bool {
					return v != 1
				},
			},
			want: []int{0, 2},
		},
		{
			name: "func([]*int(nil), true) == []*int{}",
			args: args{
				list: []*int(nil),
				fn: func(v int) bool {
					return true
				},
			},
			want: []*int{},
		},
		{
			name: "func([]*int{0, 1}, true) == []*int{0, 1}",
			args: args{
				list: []*int{
					&zero,
					&one,
				},
				fn: func(v *int) bool {
					return true
				},
			},
			want: []*int{
				&zero,
				&one,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Interface(tt.args.list, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filter.Interface() = %v, want %v", got, tt.want)
			}
		})
	}
}
