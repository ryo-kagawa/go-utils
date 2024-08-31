package arrays

import (
	"reflect"
	"strconv"
	"testing"
)

func TestReduce_int_string(t *testing.T) {
	type args struct {
		list         []int
		fn           func(accumulator string, value int, currentIndex int, array []int) string
		initialValue string
	}
	tests := []struct {
		name string
		args args
		want string
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.list, tt.args.fn, tt.args.initialValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduce_int_int(t *testing.T) {
	type args struct {
		list         []int
		fn           func(accumulator int, value int, currentIndex int, array []int) int
		initialValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.list, tt.args.fn, tt.args.initialValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}
