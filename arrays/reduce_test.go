package arrays

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_reduce_Interface(t *testing.T) {
	type args struct {
		list interface{}
		fn   interface{}
	}
	tests := []struct {
		name string
		r    reduce
		args args
		want interface{}
	}{
		{
			name: "func([]int(nil), func(acc string, cur int, i int, src []int) string == \"\"",
			args: args{
				list: []int(nil),
				fn: func(acc string, cur int, i int, src []int) string {
					return acc + strconv.FormatInt(int64(cur), 10)
				},
			},
			want: "",
		},
		{
			name: "func([]int{}, func(acc string, cur int, i int, src []int) string == \"\"",
			args: args{
				list: []int{},
				fn: func(acc string, cur int, i int, src []int) string {
					return acc + strconv.FormatInt(int64(cur), 10)
				},
			},
			want: "",
		},
		{
			name: "func([]int{1, 2}, func(acc string, cur int, i int, src []int) string == \"12\"",
			args: args{
				list: []int{1, 2},
				fn: func(acc string, cur int, i int, src []int) string {
					return acc + strconv.FormatInt(int64(cur), 10)
				},
			},
			want: "12",
		},
		{
			name: "func([]int{1, 2}, func(acc int, cur int, i int, src []int) int == 3",
			args: args{
				list: []int{1, 2},
				fn: func(acc int, cur int, i int, src []int) int {
					return acc + cur
				},
			},
			want: 3,
		},
		{
			name: "func([]int{1, 2}, func(acc int, cur int) int == 3",
			args: args{
				list: []int{1, 2},
				fn: func(acc int, cur int) int {
					return acc + cur
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Interface(tt.args.list, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reduce.Interface() = %v, want %v", got, tt.want)
			}
		})
	}
}
