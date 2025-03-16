//go:build unix || plan9 || windows

package commandline

import (
	"errors"
	"os"
	"reflect"
	"testing"
)

type validateStruct struct {
	ValidateError string `key:"ValidateError"`
}

func (v validateStruct) Validate() error {
	if v.ValidateError != "" {
		return errors.New(v.ValidateError)
	}
	return nil
}

func TestEnvironmentVariableParse_ValidateStruct(t *testing.T) {
	tests := []struct {
		preprocess func()
		name       string
		want       validateStruct
		wantErr    bool
	}{
		{
			preprocess: func() {

			},
			name: "Environment\nEnvironmentVariableParse() == validateStruct{ValidateError: \"\"}, nil",
			want: validateStruct{
				ValidateError: "",
			},
			wantErr: false,
		},
		{
			preprocess: func() {
				os.Setenv("ValidateError", "validate error")
			},
			name: "Environment\nValidateError=\"validate error\"\nEnvironmentVariableParse() == validateStruct{ValidateError: \"\"}, ",
			want: validateStruct{
				ValidateError: "validate error",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv()
			tt.preprocess()
			got, err := EnvironmentVariableParse[validateStruct]()
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
