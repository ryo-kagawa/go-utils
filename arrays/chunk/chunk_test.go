package chunk

import (
	"reflect"
	"testing"
)

func TestInt(t *testing.T) {
	type args struct {
		list      []int
		chunkSize int
	}
	tests := []struct {
		name string
		args args
		want [][]int
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
				list:      []int(nil),
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.args.list, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}
