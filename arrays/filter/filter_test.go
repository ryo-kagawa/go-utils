package filter

import (
	"reflect"
	"testing"
)

func TestInterface(t *testing.T) {
	zero := 0
	one := 1
	type args struct {
		list interface{}
		f    interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "func([]int(nil), true) == []int{}",
			args: args{
				list: []int(nil),
				f: func(v int) bool {
					return true
				},
			},
			want: []int{},
		},
		{
			name: "func([]int{}, true) == []int{}",
			args: args{
				list: []int{},
				f: func(v int) bool {
					return true
				},
			},
			want: []int{},
		},
		{
			name: "func([]int{0}, v == 0) == []int{0}",
			args: args{
				list: []int{0},
				f: func(v int) bool {
					return v == 0
				},
			},
			want: []int{0},
		},
		{
			name: "func([]int{1}, v == 0) == []int{}",
			args: args{
				list: []int{1},
				f: func(v int) bool {
					return v == 0
				},
			},
			want: []int{},
		},
		{
			name: "func([]int{0, 1, 2}, v != 1) == []int{0, 2}",
			args: args{
				list: []int{0, 1, 2},
				f: func(v int) bool {
					return v != 1
				},
			},
			want: []int{0, 2},
		},
		{
			name: "func([]*int(nil), true) == []*int{}",
			args: args{
				list: []*int(nil),
				f: func(v int) bool {
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
				f: func(v *int) bool {
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
			if got := Interface(tt.args.list, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interface() = %v, want %v", got, tt.want)
			}
		})
	}
}
