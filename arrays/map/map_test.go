package filter

import (
	"reflect"
	"strconv"
	"testing"
)

func TestInt(t *testing.T) {
	zeroString := "0"
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
			name: "func([]int(nil)) == []string{}",
			args: args{
				list: []int(nil),
				f: func(v int) string {
					return strconv.FormatInt(int64(v), 10)
				},
			},
			want: []string{},
		},
		{
			name: "func([]int{}) == []string{}",
			args: args{
				list: []int{},
				f: func(v int) string {
					return strconv.FormatInt(int64(v), 10)
				},
			},
			want: []string{},
		},
		{
			name: "func([]int{0}) == []string{\"0\"}",
			args: args{
				list: []int{0},
				f: func(v int) string {
					return strconv.FormatInt(int64(v), 10)
				},
			},
			want: []string{"0"},
		},
		{
			name: "func([]int{0, 1}) == []string{\"0\", \"1\"}",
			args: args{
				list: []int{0, 1},
				f: func(v int) string {
					return strconv.FormatInt(int64(v), 10)
				},
			},
			want: []string{"0", "1"},
		},
		{
			name: "func([]int{0}) == []*string{\"0\"}",
			args: args{
				list: []int{0},
				f: func(v int) *string {
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
				f: func(v *int) string {
					return strconv.FormatInt(int64(*v), 10)
				},
			},
			want: []string{},
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
