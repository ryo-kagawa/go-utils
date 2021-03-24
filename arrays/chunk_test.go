package arrays

import (
	"reflect"
	"testing"
)

func Test_chunk_Interface(t *testing.T) {
	type args struct {
		list      interface{}
		chunkSize int
	}
	tests := []struct {
		name string
		c    chunk
		args args
		want interface{}
	}{
		{
			name: "func([]int(nil), 1) == [][]int{}",
			args: args{
				list:      []int(nil),
				chunkSize: 1,
			},
			want: [][]int{},
		},
		{
			name: "func([]int{}, 1) == [][]int{}",
			args: args{
				list:      []int{},
				chunkSize: 1,
			},
			want: [][]int{},
		},
		{
			name: "func([]int{1}, 1) == [][]int{{1}}",
			args: args{
				list:      []int{1},
				chunkSize: 1,
			},
			want: [][]int{{1}},
		},
		{
			name: "func([]int{1, 2}, 1) ==[][]int{{1}, {2}}",
			args: args{
				list:      []int{1, 2},
				chunkSize: 1,
			},
			want: [][]int{{1}, {2}},
		},
		{
			name: "func([]int{1, 2, 3, 4, 5}, 2) == [][]int{{1, 2}, {3, 4}, {5}}",
			args: args{
				list:      []int{1, 2, 3, 4, 5},
				chunkSize: 2,
			},
			want: [][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			name: "func([]*int(nil), 1) == [][]*int{}",
			args: args{
				list:      []*int(nil),
				chunkSize: 1,
			},
			want: [][]*int{},
		},
		{
			name: "func([]*int[nil], 1) == [][]*int{nil}",
			args: args{
				list:      []*int{nil},
				chunkSize: 1,
			},
			want: [][]*int{{nil}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Interface(tt.args.list, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chunk.Interface() = %v, want %v", got, tt.want)
			}
		})
	}
}
