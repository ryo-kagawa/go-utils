package conditional

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
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
			if got := Func(tt.args.condition, tt.args.trueFunc, tt.args.falseFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Func() = %v, want %v", got, tt.want)
			}
		})
	}
}
