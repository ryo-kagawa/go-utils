package conditional

import (
	"reflect"
	"testing"
)

func TestInterface(t *testing.T) {
	type args struct {
		condition  bool
		trueValue  interface{}
		falseValue interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Interface(true, true, false) == interface{}(true)",
			args: args{
				condition:  true,
				trueValue:  true,
				falseValue: false,
			},
			want: true,
		},
		{
			name: "Interface(false, true, false) == interface{}(false)",
			args: args{
				condition:  true,
				trueValue:  false,
				falseValue: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Interface(tt.args.condition, tt.args.trueValue, tt.args.falseValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBool(t *testing.T) {
	type args struct {
		condition  bool
		trueValue  bool
		falseValue bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Bool(true, true, false) == true",
			args: args{
				condition:  true,
				trueValue:  true,
				falseValue: false,
			},
			want: true,
		},
		{
			name: "Bool(false, true, false) == false",
			args: args{
				condition:  false,
				trueValue:  true,
				falseValue: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bool(tt.args.condition, tt.args.trueValue, tt.args.falseValue); got != tt.want {
				t.Errorf("Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	type args struct {
		condition  bool
		trueValue  float32
		falseValue float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "Float32(true, 0, 1) == float32(0)",
			args: args{
				condition:  true,
				trueValue:  0,
				falseValue: 1,
			},
			want: 0,
		},
		{
			name: "Float32(false, 0, 1) == float32(1)",
			args: args{
				condition:  false,
				trueValue:  0,
				falseValue: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32(tt.args.condition, tt.args.trueValue, tt.args.falseValue); got != tt.want {
				t.Errorf("Float32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	type args struct {
		condition  bool
		trueValue  float64
		falseValue float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Float64(true, 0, 1) == float64(0)",
			args: args{
				condition:  true,
				trueValue:  0,
				falseValue: 1,
			},
			want: 0,
		},
		{
			name: "Float64(false, 0, 1) == float64(1)",
			args: args{
				condition:  false,
				trueValue:  0,
				falseValue: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64(tt.args.condition, tt.args.trueValue, tt.args.falseValue); got != tt.want {
				t.Errorf("Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt(t *testing.T) {
	type args struct {
		condition  bool
		trueValue  int
		falseValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Int(true, 0, 1) == int(0)",
			args: args{
				condition:  true,
				trueValue:  0,
				falseValue: 1,
			},
			want: 0,
		},
		{
			name: "Int(false, 0, 1) == int(1)",
			args: args{
				condition:  false,
				trueValue:  0,
				falseValue: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.args.condition, tt.args.trueValue, tt.args.falseValue); got != tt.want {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8(t *testing.T) {
	type args struct {
		condition  bool
		trueValue  int8
		falseValue int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{
			name: "Int8(true, 0, 1) == int8(0)",
			args: args{
				condition:  true,
				trueValue:  0,
				falseValue: 1,
			},
			want: 0,
		},
		{
			name: "Int8(false, 0, 1) == int8(1)",
			args: args{
				condition:  false,
				trueValue:  0,
				falseValue: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8(tt.args.condition, tt.args.trueValue, tt.args.falseValue); got != tt.want {
				t.Errorf("Int8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16(t *testing.T) {
	type args struct {
		condition  bool
		trueValue  int16
		falseValue int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			name: "Int16(true, 0, 1) == int16(0)",
			args: args{
				condition:  true,
				trueValue:  0,
				falseValue: 1,
			},
			want: 0,
		},
		{
			name: "Int16(false, 0, 1) == int16(1)",
			args: args{
				condition:  false,
				trueValue:  0,
				falseValue: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int16(tt.args.condition, tt.args.trueValue, tt.args.falseValue); got != tt.want {
				t.Errorf("Int16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	type args struct {
		condition  bool
		trueValue  int32
		falseValue int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Int32(true, 0, 1) == int32(0)",
			args: args{
				condition:  true,
				trueValue:  0,
				falseValue: 1,
			},
			want: 0,
		},
		{
			name: "Int32(false, 0, 1) == int32(1)",
			args: args{
				condition:  false,
				trueValue:  0,
				falseValue: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32(tt.args.condition, tt.args.trueValue, tt.args.falseValue); got != tt.want {
				t.Errorf("Int32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	type args struct {
		condition  bool
		trueValue  int64
		falseValue int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Int64(true, 0, 1) == int64(0)",
			args: args{
				condition:  true,
				trueValue:  0,
				falseValue: 1,
			},
			want: 0,
		},
		{
			name: "Int64(false, 0, 1) == int64(1)",
			args: args{
				condition:  false,
				trueValue:  0,
				falseValue: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64(tt.args.condition, tt.args.trueValue, tt.args.falseValue); got != tt.want {
				t.Errorf("Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	type args struct {
		condition  bool
		trueValue  string
		falseValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "String(true, \"a\", \"b\") == \"a\"",
			args: args{
				condition:  true,
				trueValue:  "a",
				falseValue: "b",
			},
			want: "a",
		},
		{
			name: "String(true, \"a\", \"b\") == \"b\"",
			args: args{
				condition:  false,
				trueValue:  "a",
				falseValue: "b",
			},
			want: "b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.condition, tt.args.trueValue, tt.args.falseValue); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUInt(t *testing.T) {
	type args struct {
		condition  bool
		trueValue  uint
		falseValue uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "UInt(true, 0, 1) == uint(0)",
			args: args{
				condition:  true,
				trueValue:  0,
				falseValue: 1,
			},
			want: 0,
		},
		{
			name: "UInt(false, 0, 1) == uint(1)",
			args: args{
				condition:  false,
				trueValue:  0,
				falseValue: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt(tt.args.condition, tt.args.trueValue, tt.args.falseValue); got != tt.want {
				t.Errorf("UInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUInt8(t *testing.T) {
	type args struct {
		condition  bool
		trueValue  uint8
		falseValue uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			name: "UInt8(true, 0, 1) == uint8(0)",
			args: args{
				condition:  true,
				trueValue:  0,
				falseValue: 1,
			},
			want: 0,
		},
		{
			name: "UInt8(false, 0, 1) == uint8(1)",
			args: args{
				condition:  false,
				trueValue:  0,
				falseValue: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt8(tt.args.condition, tt.args.trueValue, tt.args.falseValue); got != tt.want {
				t.Errorf("UInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUInt16(t *testing.T) {
	type args struct {
		condition  bool
		trueValue  uint16
		falseValue uint16
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "UInt16(true, 0, 1) == uint16(0)",
			args: args{
				condition:  true,
				trueValue:  0,
				falseValue: 1,
			},
			want: 0,
		},
		{
			name: "UInt16(false, 0, 1) == uint16(1)",
			args: args{
				condition:  false,
				trueValue:  0,
				falseValue: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt16(tt.args.condition, tt.args.trueValue, tt.args.falseValue); got != tt.want {
				t.Errorf("UInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUInt32(t *testing.T) {
	type args struct {
		condition  bool
		trueValue  uint32
		falseValue uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "UInt32(true, 0, 1) == uint32(0)",
			args: args{
				condition:  true,
				trueValue:  0,
				falseValue: 1,
			},
			want: 0,
		},
		{
			name: "UInt32(false, 0, 1) == uint32(1)",
			args: args{
				condition:  false,
				trueValue:  0,
				falseValue: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt32(tt.args.condition, tt.args.trueValue, tt.args.falseValue); got != tt.want {
				t.Errorf("UInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUInt64(t *testing.T) {
	type args struct {
		condition  bool
		trueValue  uint64
		falseValue uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "UInt64(true, 0, 1) == uint64(0)",
			args: args{
				condition:  true,
				trueValue:  0,
				falseValue: 1,
			},
			want: 0,
		},
		{
			name: "UInt64(false, 0, 1) == uint64(1)",
			args: args{
				condition:  false,
				trueValue:  0,
				falseValue: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt64(tt.args.condition, tt.args.trueValue, tt.args.falseValue); got != tt.want {
				t.Errorf("UInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
