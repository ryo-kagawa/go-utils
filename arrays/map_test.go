package arrays

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMap_int_string(t *testing.T) {
	type args struct {
		list []int
		fn   func(value int) string
	}
	tests := []struct {
		name string
		args args
		want []string
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.list, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap_int_stringpointer(t *testing.T) {
	zeroString := "0"
	type args struct {
		list []int
		fn   func(value int) *string
	}
	tests := []struct {
		name string
		args args
		want []*string
	}{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.list, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap_intpointer_string(t *testing.T) {
	type args struct {
		list []*int
		fn   func(value *int) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
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
			if got := Map(tt.args.list, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
