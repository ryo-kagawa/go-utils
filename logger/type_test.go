package logger

import (
	"reflect"
	"testing"
)

func TestParseType(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want Type
	}{
		{
			name: "NONE",
			args: args{v: "NONE"},
			want: TypeNone,
		},
		{
			name: "STDOUT",
			args: args{v: "STDOUT"},
			want: TypeStandardOut,
		},
		{
			name: "STDERR",
			args: args{v: "STDERR"},
			want: TypeStandardError,
		},
		{
			name: "FILE",
			args: args{v: "FILE"},
			want: TypeFile,
		},
		{
			name: "Unknown",
			args: args{v: ""},
			want: TypeUnknown,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseType(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseType() = %v, want %v", got, tt.want)
			}
		})
	}
}
