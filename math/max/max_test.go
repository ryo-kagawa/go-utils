package max

import (
	"testing"
)

func TestFloat64(t *testing.T) {
	type args struct {
		value     float64
		valueList []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "left",
			args: args{
				value:     20,
				valueList: []float64{10},
			},
			want: 20,
		},
		{
			name: "right",
			args: args{
				value:     10,
				valueList: []float64{20},
			},
			want: 20,
		},
		{
			name: "none",
			args: args{
				value:     10,
				valueList: []float64{},
			},
			want: 10,
		},
		{
			name: "multiValue",
			args: args{
				value:     10,
				valueList: []float64{20, 30, 25},
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64(tt.args.value, tt.args.valueList...); got != tt.want {
				t.Errorf("Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt(t *testing.T) {
	type args struct {
		value     int
		valueList []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "left",
			args: args{
				value:     20,
				valueList: []int{0},
			},
			want: 20,
		},
		{
			name: "right",
			args: args{
				value:     10,
				valueList: []int{20},
			},
			want: 20,
		},
		{
			name: "none",
			args: args{
				value:     10,
				valueList: []int{},
			},
			want: 10,
		},
		{
			name: "multiValue",
			args: args{
				value:     10,
				valueList: []int{20, 30, 25},
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.args.value, tt.args.valueList...); got != tt.want {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUInt(t *testing.T) {
	type args struct {
		value     uint
		valueList []uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "left",
			args: args{
				value:     20,
				valueList: []uint{0},
			},
			want: 20,
		},
		{
			name: "right",
			args: args{
				value:     10,
				valueList: []uint{20},
			},
			want: 20,
		},
		{
			name: "none",
			args: args{
				value:     10,
				valueList: []uint{},
			},
			want: 10,
		},
		{
			name: "multiValue",
			args: args{
				value:     10,
				valueList: []uint{20, 30, 25},
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt(tt.args.value, tt.args.valueList...); got != tt.want {
				t.Errorf("UInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
