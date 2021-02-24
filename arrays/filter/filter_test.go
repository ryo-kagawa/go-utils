package filter

import (
	"reflect"
	"testing"
)

func TestInt(t *testing.T) {
	type args struct {
		list []int
		f    func(v int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1 to 1",
			args: args{
				list: []int{0},
				f: func(v int) bool {
					return v == 0
				},
			},
			want: []int{0},
		},
		{
			name: "2 to 1",
			args: args{
				list: []int{0, 1},
				f: func(v int) bool {
					return v == 1
				},
			},
			want: []int{1},
		},
		{
			name: "1 to 0",
			args: args{
				list: []int{0},
				f: func(v int) bool {
					return false
				},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.args.list, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}
