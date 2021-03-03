package find

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
	type wants struct {
		want1 interface{}
		want2 bool
	}
	tests := []struct {
		name  string
		args  args
		wants wants
	}{
		{
			name: "func([]int(nil), true) == 0, false",
			args: args{
				list: []int(nil),
				f: func(v int) bool {
					return true
				},
			},
			wants: wants{
				want1: 0,
				want2: false,
			},
		},
		{
			name: "func([]int{}, true) == 0, false",
			args: args{
				list: []int{},
				f: func(v int) bool {
					return true
				},
			},
			wants: wants{
				0,
				false,
			},
		},
		{
			name: "func([]int{0}, v == 0) == 0, true",
			args: args{
				list: []int{0},
				f: func(v int) bool {
					return v == 0
				},
			},
			wants: wants{
				0,
				true,
			},
		},
		{
			name: "func([]int{1}, v == 0) == 0, false",
			args: args{
				list: []int{1},
				f: func(v int) bool {
					return v == 0
				},
			},
			wants: wants{
				want1: 0,
				want2: false,
			},
		},
		{
			name: "func([]int{0, 1}, v == 1) == 1, true",
			args: args{
				list: []int{0, 1},
				f: func(v int) bool {
					return v == 1
				},
			},
			wants: wants{
				1,
				true,
			},
		},
		{
			name: "func([]int{0, 1, 2}, v != 1) == 0, true",
			args: args{
				list: []int{0, 1, 2},
				f: func(v int) bool {
					return v != 1
				},
			},
			wants: wants{
				want1: 0,
				want2: true,
			},
		},
		{
			name: "func([]*int(nil), true) == []*int{}",
			args: args{
				list: []*int(nil),
				f: func(v int) bool {
					return true
				},
			},
			wants: wants{
				want1: (*int)(nil),
				want2: false,
			},
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
			wants: wants{
				want1: &zero,
				want2: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := Interface(tt.args.list, tt.args.f)
			if !reflect.DeepEqual(got1, tt.wants.want1) || got2 != tt.wants.want2 {
				t.Errorf("Interface() got1 = %v, want1 %v, got2 = %v, want2 = %v", got1, tt.wants.want1, got2, tt.wants.want2)
			}
		})
	}
}
