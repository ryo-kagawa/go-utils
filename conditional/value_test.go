package conditional

import (
	"testing"
)

func TestValue(t *testing.T) {
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
			if got := Value(tt.args.condition, tt.args.trueValue, tt.args.falseValue); got != tt.want {
				t.Errorf("Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}
