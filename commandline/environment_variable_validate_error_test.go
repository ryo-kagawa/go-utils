package commandline

import (
	"os"
	"reflect"
	"testing"
)

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
