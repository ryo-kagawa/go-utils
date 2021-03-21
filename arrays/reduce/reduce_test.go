package filter

import (
	"reflect"
	"strconv"
	"testing"
)

func TestInterface(t *testing.T) {
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
			name: "func([]int(nil), func(acc string, cur int, i int, src []int) string == \"\"",
			args: args{
				list: []int(nil),
				f: func(acc string, cur int, i int, src []int) string {
					return acc + strconv.FormatInt(int64(cur), 10)
				},
			},
			want: "",
		},
		{
			name: "func([]int{}, func(acc string, cur int, i int, src []int) string == \"\"",
			args: args{
				list: []int{},
				f: func(acc string, cur int, i int, src []int) string {
					return acc + strconv.FormatInt(int64(cur), 10)
				},
			},
			want: "",
		},
		{
			name: "func([]int{1, 2}, func(acc string, cur int, i int, src []int) string == \"12\"",
			args: args{
				list: []int{1, 2},
				f: func(acc string, cur int, i int, src []int) string {
					return acc + strconv.FormatInt(int64(cur), 10)
				},
			},
			want: "12",
		},
		{
			name: "func([]int{1, 2}, func(acc int, cur int, i int, src []int) int == 3",
			args: args{
				list: []int{1, 2},
				f: func(acc int, cur int, i int, src []int) int {
					return acc + cur
				},
			},
			want: 3,
		},
		{
			name: "func([]int{1, 2}, func(acc int, cur int) int == 3",
			args: args{
				list: []int{1, 2},
				f: func(acc int, cur int) int {
					return acc + cur
				},
			},
			want: 3,
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
