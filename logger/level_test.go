package logger

import (
	"reflect"
	"testing"
)

func TestParseLevel(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want Level
	}{
		{
			name: "FATAL",
			args: args{v: "FATAL"},
			want: LevelFatal,
		},
		{
			name: "ERROR",
			args: args{v: "ERROR"},
			want: LevelError,
		},
		{
			name: "WARN",
			args: args{v: "WARN"},
			want: LevelWarn,
		},
		{
			name: "INFO",
			args: args{v: "INFO"},
			want: LevelInfo,
		},
		{
			name: "DEBUG",
			args: args{v: "DEBUG"},
			want: LevelDebug,
		},
		{
			name: "Unknown",
			args: args{v: ""},
			want: LevelUnknown,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseLevel(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
