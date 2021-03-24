package arrays

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_mapImpl_Interface(t *testing.T) {
	zeroString := "0"
	type args struct {
		list interface{}
		fn   interface{}
	}
	tests := []struct {
		name string
		m    mapImpl
		args args
		want interface{}
	}{
		{
			name: "func([]int(nil)) == []string{}",
			args: args{
				list: []int(nil),
				fn: func(v int) string {
					return strconv.FormatInt(int64(v), 10)
				},
			},
			want: []string{},
		},
		{
			name: "func([]int{}) == []string{}",
			args: args{
				list: []int{},
				fn: func(v int) string {
					return strconv.FormatInt(int64(v), 10)
				},
			},
			want: []string{},
		},
		{
			name: "func([]int{0}) == []string{\"0\"}",
			args: args{
				list: []int{0},
				fn: func(v int) string {
					return strconv.FormatInt(int64(v), 10)
				},
			},
			want: []string{"0"},
		},
		{
			name: "func([]int{0, 1}) == []string{\"0\", \"1\"}",
			args: args{
				list: []int{0, 1},
				fn: func(v int) string {
					return strconv.FormatInt(int64(v), 10)
				},
			},
			want: []string{"0", "1"},
		},
		{
			name: "func([]int{0}) == []*string{\"0\"}",
			args: args{
				list: []int{0},
				fn: func(v int) *string {
					x := strconv.FormatInt(int64(v), 10)
					return &x
				},
			},
			want: []*string{&zeroString},
		},
		{
			name: "func([]*int(nil)) == []string{}",
			args: args{
				list: []*int(nil),
				fn: func(v *int) string {
					return strconv.FormatInt(int64(*v), 10)
				},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Interface(tt.args.list, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapImpl.Interface() = %v, want %v", got, tt.want)
			}
		})
	}
}
