package arrays

import (
	"reflect"
	"testing"
)

func TestFind_int(t *testing.T) {
	type args struct {
		list []int
		fn   func(value int) bool
	}
	tests := []struct {
		name  string
		args  args
		want1 int
		want2 bool
	}{
		{
			name: "func([]int(nil), true) == 0, false",
			args: args{
				list: []int(nil),
				fn: func(v int) bool {
					return true
				},
			},
			want1: 0,
			want2: false,
		},
		{
			name: "func([]int{}, true) == 0, false",
			args: args{
				list: []int{},
				fn: func(v int) bool {
					return true
				},
			},
			want1: 0,
			want2: false,
		},
		{
			name: "func([]int{0}, v == 0) == 0, true",
			args: args{
				list: []int{0},
				fn: func(v int) bool {
					return v == 0
				},
			},
			want1: 0,
			want2: true,
		},
		{
			name: "func([]int{1}, v == 0) == 0, false",
			args: args{
				list: []int{1},
				fn: func(v int) bool {
					return v == 0
				},
			},
			want1: 0,
			want2: false,
		},
		{
			name: "func([]int{0, 1}, v == 1) == 1, true",
			args: args{
				list: []int{0, 1},
				fn: func(v int) bool {
					return v == 1
				},
			},
			want1: 1,
			want2: true,
		},
		{
			name: "func([]int{0, 1, 2}, v != 1) == 0, true",
			args: args{
				list: []int{0, 1, 2},
				fn: func(v int) bool {
					return v != 1
				},
			},
			want1: 0,
			want2: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := Find(tt.args.list, tt.args.fn)
			if got1 != tt.want1 {
				t.Errorf("Find() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Find() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestFind_intpointer(t *testing.T) {
	zero := 0
	one := 1
	type args struct {
		list []*int
		fn   func(value *int) bool
	}
	tests := []struct {
		name  string
		args  args
		want1 *int
		want2 bool
	}{
		{
			name: "func([]*int(nil), true) == []*int{}",
			args: args{
				list: []*int(nil),
				fn: func(v *int) bool {
					return true
				},
			},
			want1: (*int)(nil),
			want2: false,
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
			want1: &zero,
			want2: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := Find(tt.args.list, tt.args.fn)
			if got1 != tt.want1 {
				t.Errorf("Find() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Find() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
