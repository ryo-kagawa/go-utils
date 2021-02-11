package logger

import (
	"bytes"
	"reflect"
	"testing"
)

func TestNewLogConsole(t *testing.T) {
	type args struct {
		level Level
	}
	tests := []struct {
		name       string
		args       args
		want       Log
		wantWriter string
	}{
		{
			name: "",
			args: args{
				level: LevelFatal,
			},
			want: &LogConsole{
				Level:  LevelFatal,
				Writer: &bytes.Buffer{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			if got := NewLogConsole(tt.args.level, writer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLog() = %v, want %v", got, tt.want)
			}
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("NewLog() = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}

func TestLogConsole_IsLog(t *testing.T) {
	type args struct {
		level Level
	}
	type test struct {
		args args
		want bool
	}
	createTest := func(level Level, want bool) test {
		return test{
			args: args{
				level: level,
			},
			want: want,
		}
	}
	tests := []struct {
		name string
		l    *LogConsole
		test []test
	}{
		{
			name: "FATAL",
			l:    NewLogConsole(LevelFatal, &bytes.Buffer{}),
			test: []test{
				createTest(LevelFatal, true),
				createTest(LevelError, false),
				createTest(LevelWarn, false),
				createTest(LevelInfo, false),
				createTest(LevelDebug, false),
			},
		},
		{
			name: "ERROR",
			l:    NewLogConsole(LevelError, &bytes.Buffer{}),
			test: []test{
				createTest(LevelFatal, true),
				createTest(LevelError, true),
				createTest(LevelWarn, false),
				createTest(LevelInfo, false),
				createTest(LevelDebug, false),
			},
		},
		{
			name: "WARN",
			l:    NewLogConsole(LevelWarn, &bytes.Buffer{}),
			test: []test{
				createTest(LevelFatal, true),
				createTest(LevelError, true),
				createTest(LevelWarn, true),
				createTest(LevelInfo, false),
				createTest(LevelDebug, false),
			},
		},
		{
			name: "INFO",
			l:    NewLogConsole(LevelInfo, &bytes.Buffer{}),
			test: []test{
				createTest(LevelFatal, true),
				createTest(LevelError, true),
				createTest(LevelWarn, true),
				createTest(LevelInfo, true),
				createTest(LevelDebug, false),
			},
		},
		{
			name: "DEBUG",
			l:    NewLogConsole(LevelDebug, &bytes.Buffer{}),
			test: []test{
				createTest(LevelFatal, true),
				createTest(LevelError, true),
				createTest(LevelWarn, true),
				createTest(LevelInfo, true),
				createTest(LevelDebug, true),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, x := range tt.test {
				if got := tt.l.IsLog(x.args.level); got != x.want {
					t.Errorf("LogConsole.IsLog(%v) = %v, want %v", x.args.level, got, x.want)
				}
			}
		})
	}
}

func TestLogConsole_Log(t *testing.T) {
	buf := &bytes.Buffer{}
	log := NewLogConsole(LevelFatal, buf)
	value := "a"
	expected := "a"
	log.Log(value)
	actual := buf.String()
	if actual != expected {
		t.Errorf("LogConsole.Log(%v), actual %v, expected %v", value, actual, expected)
	}
}
