//go:build unix

package commandline

import (
	"os"
	"reflect"
	"testing"
)

func TestEnvironmentVariableParse_normalStruct(t *testing.T) {
	type normalStruct struct {
		BoolValue   bool   `key:"BoolValue"`
		boolValue   bool   `key:"boolValue"`
		IntValue    int    `key:"IntValue"`
		intValue    int    `key:"intValue"`
		StringValue string `key:"StringValue"`
		stringValue string `key:"stringValue"`
	}

	tests := []struct {
		preprocess func()
		name       string
		want       normalStruct
		wantErr    bool
	}{
		{
			preprocess: func() {
				os.Setenv("BoolValue", "")
			},
			name: "Environment\nBoolValue\nEnvironmentVariableParse() == normalStruct{BoolValue: true}, nil",
			want: normalStruct{
				BoolValue: true,
			},
			wantErr: false,
		},
		{
			preprocess: func() {
				os.Setenv("boolValue", "")
			},
			name: "Environment\nboolValue\nEnvironmentVariableParse() == normalStruct{boolValue: true}, nil",
			want: normalStruct{
				boolValue: true,
			},
			wantErr: false,
		},
		{
			preprocess: func() {
				os.Setenv("IntValue", "1")
			},
			name: "Environment\nIntValue=1\nEnvironmentVariableParse() == normalStruct{IntValue: 1}, nil",
			want: normalStruct{
				IntValue: 1,
			},
			wantErr: false,
		},
		{
			preprocess: func() {
				os.Setenv("intValue", "1")
			},
			name: "Environment\nintValue=1\nEnvironmentVariableParse() == normalStruct{intValue: 1}, nil",
			want: normalStruct{
				intValue: 1,
			},
			wantErr: false,
		},
		{
			preprocess: func() {
				os.Setenv("StringValue", "a")
			},
			name: "Environment\nStringValue=a\nEnvironmentVariableParse() == normalStruct{StringValue: a}, nil",
			want: normalStruct{
				StringValue: "a",
			},
			wantErr: false,
		},
		{
			preprocess: func() {
				os.Setenv("stringValue", "a")
			},
			name: "Environment\nstringValue=a\nEnvironmentVariableParse() == normalStruct{stringValue: a}, nil",
			want: normalStruct{
				stringValue: "a",
			},
			wantErr: false,
		},
		{
			preprocess: func() {
				os.Setenv("BoolValue", "a")
			},
			name:    "Environment\nBoolValue=a\nEnvironmentVariableParse() == normalStruct{}, error",
			want:    normalStruct{},
			wantErr: true,
		},
		{
			preprocess: func() {
				os.Setenv("IntValue", "")
			},
			name:    "Environment\nIntValue=\nEnvironmentVariableParse() == normalStruct{}, error",
			want:    normalStruct{},
			wantErr: true,
		},
		{
			preprocess: func() {
				os.Setenv("IntValue", "a")
			},
			name:    "Environment\nIntValue=a\nEnvironmentVariableParse() == normalStruct{}, error",
			want:    normalStruct{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv()
			tt.preprocess()
			got, err := EnvironmentVariableParse[normalStruct]()
			if (err != nil) != tt.wantErr {
				t.Errorf("EnvironmentVariableParse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EnvironmentVariableParse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvironmentVariableParse_defaultStruct(t *testing.T) {
	type defaultStruct struct {
		IntValue    int    `key:"IntValue" default:"1"`
		intValue    int    `key:"intValue" default:"1"`
		StringValue string `key:"StringValue" default:"a"`
		stringValue string `key:"stringValue" default:"a"`
	}

	tests := []struct {
		preprocess func()
		name       string
		want       defaultStruct
		wantErr    bool
	}{
		{
			preprocess: func() {},
			name:       "EnvironmentVariableParse() == defaultStruct{IntValue: 1, intValue: 1, StringValue: \"a\", stringValue: \"a\"}, nil",
			want: defaultStruct{
				IntValue:    1,
				intValue:    1,
				StringValue: "a",
				stringValue: "a",
			},
			wantErr: false,
		},
		{
			preprocess: func() {
				os.Setenv("IntValue", "0")
			},
			name: "Environment\nIntValue=0\nEnvironmentVariableParse() == defaultStruct{IntValue: 0, intValue: 1, StringValue: \"a\", stringValue: \"a\"}, error",
			want: defaultStruct{
				IntValue:    0,
				intValue:    1,
				StringValue: "a",
				stringValue: "a",
			},
			wantErr: false,
		},
		{
			preprocess: func() {
				os.Setenv("intValue", "0")
			},
			name: "Environment\nintValue=0\nEnvironmentVariableParse() == defaultStruct{IntValue: 1, intValue: 0, StringValue: \"a\", stringValue: \"a\"}, error",
			want: defaultStruct{
				IntValue:    1,
				intValue:    0,
				StringValue: "a",
				stringValue: "a",
			},
			wantErr: false,
		},
		{
			preprocess: func() {
				os.Setenv("StringValue", "b")
			},
			name: "Environment\nStringValue=b\nEnvironmentVariableParse() == defaultStruct{IntValue: 1, intValue: 1, StringValue: \"b\", stringValue: \"a\"}, error",
			want: defaultStruct{
				IntValue:    1,
				intValue:    1,
				StringValue: "b",
				stringValue: "a",
			},
			wantErr: false,
		},
		{
			preprocess: func() {
				os.Setenv("stringValue", "b")
			},
			name: "Environment\nstringValue=b\nEnvironmentVariableParse() == defaultStruct{IntValue: 1, intValue: 1, StringValue: \"a\", stringValue: \"b\"}, error",
			want: defaultStruct{
				IntValue:    1,
				intValue:    1,
				StringValue: "a",
				stringValue: "b",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv()
			tt.preprocess()
			got, err := EnvironmentVariableParse[defaultStruct]()
			if (err != nil) != tt.wantErr {
				t.Errorf("EnvironmentVariableParse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EnvironmentVariableParse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvironmentVariableParse_invalidType(t *testing.T) {
	type invalidTypeStruct struct {
		Float64Value float64 `key:"Float64Value"`
	}

	tests := []struct {
		preprocess func()
		name       string
		want       invalidTypeStruct
		wantErr    bool
	}{
		{
			preprocess: func() {
				os.Setenv("Float64Value", "1")
			},
			name:    "EnvironmentVariableParse() == invalidTypeStruct{}, error",
			want:    invalidTypeStruct{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv()
			tt.preprocess()
			got, err := EnvironmentVariableParse[invalidTypeStruct]()
			if (err != nil) != tt.wantErr {
				t.Errorf("EnvironmentVariableParse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EnvironmentVariableParse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvironmentVariableParse_invalidDefaultValue(t *testing.T) {
	type invalidDefaultStruct struct {
		IntValue int `key:"IntValue" default:""`
	}

	tests := []struct {
		preprocess func()
		name       string
		want       invalidDefaultStruct
		wantErr    bool
	}{
		{
			preprocess: func() {},
			name:       "EnvironmentVariableParse() == invalidDefaultStruct{}, error",
			want:       invalidDefaultStruct{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv()
			tt.preprocess()
			got, err := EnvironmentVariableParse[invalidDefaultStruct]()
			if (err != nil) != tt.wantErr {
				t.Errorf("EnvironmentVariableParse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EnvironmentVariableParse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvironmentVariableParse_invalidDefaultType(t *testing.T) {
	type invalidDefaultTypeStruct struct {
		Float64Value float64 `key:"IntValue" default:""`
	}

	tests := []struct {
		preprocess func()
		name       string
		want       invalidDefaultTypeStruct
		wantErr    bool
	}{
		{
			preprocess: func() {},
			name:       "EnvironmentVariableParse() == invalidDefaultTypeStruct{}, error",
			want:       invalidDefaultTypeStruct{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv()
			tt.preprocess()
			got, err := EnvironmentVariableParse[invalidDefaultTypeStruct]()
			if (err != nil) != tt.wantErr {
				t.Errorf("EnvironmentVariableParse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EnvironmentVariableParse() = %v, want %v", got, tt.want)
			}
		})
	}
}
