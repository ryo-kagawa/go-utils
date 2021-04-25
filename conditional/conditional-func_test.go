package conditional

import (
	"reflect"
	"testing"
)

func TestInterfaceFunc(t *testing.T) {
	trueValue := func() interface{} { return true }
	falseValue := func() interface{} { return false }
	type args struct {
		condition bool
		trueFunc  func() interface{}
		falseFunc func() interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "InterfaceFunc(true, trueValue, falseValue) == trueValue()",
			args: args{
				condition: true,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: trueValue(),
		},
		{
			name: "InterfaceFunc(false, trueValue, falseValue) == falseValue()",
			args: args{
				condition: false,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: falseValue(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterfaceFunc(tt.args.condition, tt.args.trueFunc, tt.args.falseFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InterfaceFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolFunc(t *testing.T) {
	trueValue := func() bool { return true }
	falseValue := func() bool { return false }
	type args struct {
		condition bool
		trueFunc  func() bool
		falseFunc func() bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "BoolFunc(true, trueValue, falseValue) == trueValue()",
			args: args{
				condition: true,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: trueValue(),
		},
		{
			name: "BoolFunc(false, trueValue, falseValue) == falseValue()",
			args: args{
				condition: false,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: falseValue(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoolFunc(tt.args.condition, tt.args.trueFunc, tt.args.falseFunc); got != tt.want {
				t.Errorf("BoolFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32Func(t *testing.T) {
	trueValue := func() float32 { return 0 }
	falseValue := func() float32 { return 1 }
	type args struct {
		condition bool
		trueFunc  func() float32
		falseFunc func() float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "Float32Func(true, trueValue, falseValue) == trueValue()",
			args: args{
				condition: true,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: trueValue(),
		},
		{
			name: "Float32Func(false, trueValue, falseValue) == falseValue()",
			args: args{
				condition: false,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: falseValue(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32Func(tt.args.condition, tt.args.trueFunc, tt.args.falseFunc); got != tt.want {
				t.Errorf("Float32Func() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64Func(t *testing.T) {
	trueValue := func() float64 { return 0 }
	falseValue := func() float64 { return 1 }
	type args struct {
		condition bool
		trueFunc  func() float64
		falseFunc func() float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Float64Func(true, trueValue, falseValue) == trueValue()",
			args: args{
				condition: true,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: trueValue(),
		},
		{
			name: "Float64Func(false, trueValue, falseValue) == falseValue()",
			args: args{
				condition: false,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: falseValue(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64Func(tt.args.condition, tt.args.trueFunc, tt.args.falseFunc); got != tt.want {
				t.Errorf("Float64Func() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntFunc(t *testing.T) {
	trueValue := func() int { return 0 }
	falseValue := func() int { return 1 }
	type args struct {
		condition bool
		trueFunc  func() int
		falseFunc func() int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "IntFunc(true, trueValue, falseValue) == trueValue()",
			args: args{
				condition: true,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: trueValue(),
		},
		{
			name: "IntFunc(false, trueValue, falseValue) == falseValue()",
			args: args{
				condition: false,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: falseValue(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntFunc(tt.args.condition, tt.args.trueFunc, tt.args.falseFunc); got != tt.want {
				t.Errorf("IntFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8Func(t *testing.T) {
	trueValue := func() int8 { return 0 }
	falseValue := func() int8 { return 1 }
	type args struct {
		condition bool
		trueFunc  func() int8
		falseFunc func() int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{
			name: "Int8Func(true, trueValue, falseValue) == trueValue()",
			args: args{
				condition: true,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: trueValue(),
		},
		{
			name: "Int8Func(false, trueValue, falseValue) == falseValue()",
			args: args{
				condition: false,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: falseValue(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8Func(tt.args.condition, tt.args.trueFunc, tt.args.falseFunc); got != tt.want {
				t.Errorf("IntFunc8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16Func(t *testing.T) {
	trueValue := func() int16 { return 0 }
	falseValue := func() int16 { return 1 }
	type args struct {
		condition bool
		trueFunc  func() int16
		falseFunc func() int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			name: "Int16Func(true, trueValue, falseValue) == trueValue()",
			args: args{
				condition: true,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: trueValue(),
		},
		{
			name: "Int16Func(false, trueValue, falseValue) == falseValue()",
			args: args{
				condition: false,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: falseValue(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int16Func(tt.args.condition, tt.args.trueFunc, tt.args.falseFunc); got != tt.want {
				t.Errorf("Int16Func() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32Func(t *testing.T) {
	trueValue := func() int32 { return 0 }
	falseValue := func() int32 { return 1 }
	type args struct {
		condition bool
		trueFunc  func() int32
		falseFunc func() int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Int32Func(true, trueValue, falseValue) == trueValue()",
			args: args{
				condition: true,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: trueValue(),
		},
		{
			name: "Int32Func(false, trueValue, falseValue) == falseValue()",
			args: args{
				condition: false,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: falseValue(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32Func(tt.args.condition, tt.args.trueFunc, tt.args.falseFunc); got != tt.want {
				t.Errorf("Int32Func() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64Func(t *testing.T) {
	trueValue := func() int64 { return 0 }
	falseValue := func() int64 { return 1 }
	type args struct {
		condition bool
		trueFunc  func() int64
		falseFunc func() int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Int64Func(true, trueValue, falseValue) == trueValue()",
			args: args{
				condition: true,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: trueValue(),
		},
		{
			name: "Int64Func(false, trueValue, falseValue) == falseValue()",
			args: args{
				condition: false,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: falseValue(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64Func(tt.args.condition, tt.args.trueFunc, tt.args.falseFunc); got != tt.want {
				t.Errorf("Int64Func() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringFunc(t *testing.T) {
	trueValue := func() string { return "a" }
	falseValue := func() string { return "b" }
	type args struct {
		condition bool
		trueFunc  func() string
		falseFunc func() string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "StringFunc(true, trueValue, falseValue) == trueValue()",
			args: args{
				condition: true,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: trueValue(),
		},
		{
			name: "StringFunc(false, trueValue, falseValue) == falseValue()",
			args: args{
				condition: false,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: falseValue(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringFunc(tt.args.condition, tt.args.trueFunc, tt.args.falseFunc); got != tt.want {
				t.Errorf("StringFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUIntFunc(t *testing.T) {
	trueValue := func() uint { return 0 }
	falseValue := func() uint { return 1 }
	type args struct {
		condition bool
		trueFunc  func() uint
		falseFunc func() uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "UIntFunc(true, trueValue, falseValue) == trueValue()",
			args: args{
				condition: true,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: trueValue(),
		},
		{
			name: "UIntFunc(false, trueValue, falseValue) == falseValue()",
			args: args{
				condition: false,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: falseValue(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UIntFunc(tt.args.condition, tt.args.trueFunc, tt.args.falseFunc); got != tt.want {
				t.Errorf("UIntFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUInt8Func(t *testing.T) {
	trueValue := func() uint8 { return 0 }
	falseValue := func() uint8 { return 1 }
	type args struct {
		condition bool
		trueFunc  func() uint8
		falseFunc func() uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			name: "UInt8Func(true, trueValue, falseValue) == trueValue()",
			args: args{
				condition: true,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: trueValue(),
		},
		{
			name: "UInt8Func(false, trueValue, falseValue) == falseValue()",
			args: args{
				condition: false,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: falseValue(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt8Func(tt.args.condition, tt.args.trueFunc, tt.args.falseFunc); got != tt.want {
				t.Errorf("UIntFunc8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUInt16Func(t *testing.T) {
	trueValue := func() uint16 { return 0 }
	falseValue := func() uint16 { return 1 }
	type args struct {
		condition bool
		trueFunc  func() uint16
		falseFunc func() uint16
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "UInt16Func(true, trueValue, falseValue) == trueValue()",
			args: args{
				condition: true,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: trueValue(),
		},
		{
			name: "UInt16Func(false, trueValue, falseValue) == falseValue()",
			args: args{
				condition: false,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: falseValue(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt16Func(tt.args.condition, tt.args.trueFunc, tt.args.falseFunc); got != tt.want {
				t.Errorf("UInt16Func() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUInt32Func(t *testing.T) {
	trueValue := func() uint32 { return 0 }
	falseValue := func() uint32 { return 1 }
	type args struct {
		condition bool
		trueFunc  func() uint32
		falseFunc func() uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "UInt32Func(true, trueValue, falseValue) == trueValue()",
			args: args{
				condition: true,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: trueValue(),
		},
		{
			name: "UInt32Func(false, trueValue, falseValue) == falseValue()",
			args: args{
				condition: false,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: falseValue(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt32Func(tt.args.condition, tt.args.trueFunc, tt.args.falseFunc); got != tt.want {
				t.Errorf("UInt32Func() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUInt64Func(t *testing.T) {
	trueValue := func() uint64 { return 0 }
	falseValue := func() uint64 { return 1 }
	type args struct {
		condition bool
		trueFunc  func() uint64
		falseFunc func() uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "UInt64Func(true, trueValue, falseValue) == trueValue()",
			args: args{
				condition: true,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: trueValue(),
		},
		{
			name: "UInt64Func(false, trueValue, falseValue) == falseValue()",
			args: args{
				condition: false,
				trueFunc:  trueValue,
				falseFunc: falseValue,
			},
			want: falseValue(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt64Func(tt.args.condition, tt.args.trueFunc, tt.args.falseFunc); got != tt.want {
				t.Errorf("UInt64Func() = %v, want %v", got, tt.want)
			}
		})
	}
}
