package commandline

import (
	"reflect"
	"testing"
)

func TestArgumentsParse_normalStruct(t *testing.T) {
	type normalStruct struct {
		BoolValue        bool     `key:"BoolValue"`
		boolValue        bool     `key:"boolValue"`
		IntValue         int      `key:"IntValue"`
		intValue         int      `key:"intValue"`
		SliceIntValue    []int    `key:"SliceIntValue"`
		sliceIntValue    []int    `key:"sliceIntValue"`
		SliceStringValue []string `key:"SliceStringValue"`
		sliceStringValue []string `key:"sliceStringValue"`
		StringValue      string   `key:"StringValue"`
		stringValue      string   `key:"stringValue"`
	}
	type args struct {
		arguments []string
	}
	tests := []struct {
		name    string
		args    args
		want    normalStruct
		wantErr bool
	}{
		{
			name: "ArgumentParse([]string{\"BoolValue\"}) == normalStruct{BoolValue: true}, nil",
			args: args{
				arguments: []string{
					"BoolValue",
				},
			},
			want: normalStruct{
				BoolValue: true,
			},
			wantErr: false,
		},
		{
			name: "ArgumentParse([]string{\"boolValue\"}) == normalStruct{boolValue: true}, nil",
			args: args{
				arguments: []string{
					"boolValue",
				},
			},
			want: normalStruct{
				boolValue: true,
			},
			wantErr: false,
		},
		{
			name: "ArgumentParse([]string{\"IntValue=1\"}) == normalStruct{IntValue: 1}, nil",
			args: args{
				arguments: []string{
					"IntValue=1",
				},
			},
			want: normalStruct{
				IntValue: 1,
			},
			wantErr: false,
		},
		{
			name: "ArgumentParse([]string{\"intValue=1\"}) == normalStruct{intValue: 1}, nil",
			args: args{
				arguments: []string{
					"intValue=1",
				},
			},
			want: normalStruct{
				intValue: 1,
			},
			wantErr: false,
		},
		{
			name: "ArgumentParse([]string{\"SliceIntValue=0\"}) == normalStruct{SliceIntValue: []int{0}}, nil",
			args: args{
				arguments: []string{
					"SliceIntValue=0",
				},
			},
			want: normalStruct{
				SliceIntValue: []int{0},
			},
			wantErr: false,
		},
		{
			name: "ArgumentParse([]string{\"SliceIntValue=0\",\"SliceIntValue=1\"}) == normalStruct{SliceIntValue: []int{0, 1}}, nil",
			args: args{
				arguments: []string{
					"SliceIntValue=0",
					"SliceIntValue=1",
				},
			},
			want: normalStruct{
				SliceIntValue: []int{0, 1},
			},
			wantErr: false,
		},
		{
			name: "ArgumentParse([]string{\"sliceIntValue=0\"}) == normalStruct{sliceIntValue: []int{0}}, nil",
			args: args{
				arguments: []string{
					"sliceIntValue=0",
				},
			},
			want: normalStruct{
				sliceIntValue: []int{0},
			},
			wantErr: false,
		},
		{
			name: "ArgumentParse([]string{\"sliceIntValue=0\",\"sliceIntValue=1\"}) == normalStruct{sliceIntValue: []int{0, 1}}, nil",
			args: args{
				arguments: []string{
					"sliceIntValue=0",
					"sliceIntValue=1",
				},
			},
			want: normalStruct{
				sliceIntValue: []int{0, 1},
			},
			wantErr: false,
		},
		{
			name: "ArgumentParse([]string{\"SliceStringValue=a\"}) == normalStruct{SliceStringValue: []string{\"a\"}}, nil",
			args: args{
				arguments: []string{
					"SliceStringValue=a",
				},
			},
			want: normalStruct{
				SliceStringValue: []string{"a"},
			},
			wantErr: false,
		},
		{
			name: "ArgumentParse([]string{\"SliceStringValue=a, SliceStringValue=b\"}) == normalStruct{SliceStringValue: []string{\"a\", \"b\"}}, nil",
			args: args{
				arguments: []string{
					"SliceStringValue=a",
					"SliceStringValue=b",
				},
			},
			want: normalStruct{
				SliceStringValue: []string{"a", "b"},
			},
			wantErr: false,
		},
		{
			name: "ArgumentParse([]string{\"sliceStringValue=a\"}) == normalStruct{sliceStringValue: []string{\"a\"}}, nil",
			args: args{
				arguments: []string{
					"sliceStringValue=a",
				},
			},
			want: normalStruct{
				sliceStringValue: []string{"a"},
			},
			wantErr: false,
		},
		{
			name: "ArgumentParse([]string{\"sliceStringValue=a, sliceStringValue=b\"}) == normalStruct{sliceStringValue: []string{\"a\", \"b\"}}, nil",
			args: args{
				arguments: []string{
					"sliceStringValue=a",
					"sliceStringValue=b",
				},
			},
			want: normalStruct{
				sliceStringValue: []string{"a", "b"},
			},
			wantErr: false,
		},
		{
			name: "ArgumentParse([]string{\"StringValue=a\"}) == normalStruct{StringValue: \"a\"}, nil",
			args: args{
				arguments: []string{
					"StringValue=a",
				},
			},
			want: normalStruct{
				StringValue: "a",
			},
			wantErr: false,
		},
		{
			name: "ArgumentParse([]string{\"stringValue=a\"}) == normalStruct{stringValue: \"a\"}, nil",
			args: args{
				arguments: []string{
					"stringValue=a",
				},
			},
			want: normalStruct{
				stringValue: "a",
			},
			wantErr: false,
		},
		{
			name: "ArgumentParse([]string{\"BoolValue=\"}) == normalStruct{}, error",
			args: args{
				arguments: []string{
					"BoolValue=",
				},
			},
			want:    normalStruct{},
			wantErr: true,
		},
		{
			name: "ArgumentParse([]string{\"IntValue\"}) == normalStruct{}, error",
			args: args{
				arguments: []string{
					"IntValue",
				},
			},
			want:    normalStruct{},
			wantErr: true,
		},
		{
			name: "ArgumentParse([]string{\"IntValue=a\"}) == normalStruct{}, error",
			args: args{
				arguments: []string{
					"IntValue=a",
				},
			},
			want:    normalStruct{},
			wantErr: true,
		},
		{
			name: "ArgumentParse([]string{\"SliceIntValue=a\"}) == normalStruct{}, error",
			args: args{
				arguments: []string{
					"SliceIntValue=a",
				},
			},
			want:    normalStruct{},
			wantErr: true,
		},
		{
			name: "ArgumentParse([]string{\"a=a\"}) == normalStruct{}, error",
			args: args{
				arguments: []string{
					"a=a",
				},
			},
			want:    normalStruct{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ArgumentsParse[normalStruct](tt.args.arguments)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArgumentParse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArgumentParse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArgumentsParse_defaultStruct(t *testing.T) {
	type defaultStruct struct {
		IntValue    int    `key:"IntValue" default:"1"`
		intValue    int    `key:"intValue" default:"1"`
		StringValue string `key:"StringValue" default:"a"`
		stringValue string `key:"stringValue" default:"a"`
	}
	type args struct {
		arguments []string
	}
	tests := []struct {
		name    string
		args    args
		want    defaultStruct
		wantErr bool
	}{
		{
			name: "ArgumentParse([]string{}) == defaultStruct{IntValue: 1, intValue: 1, StringValue: \"a\", stringValue: \"a\"}, nil",
			args: args{
				arguments: []string{},
			},
			want: defaultStruct{
				IntValue:    1,
				intValue:    1,
				StringValue: "a",
				stringValue: "a",
			},
			wantErr: false,
		},
		{
			name: "ArgumentParse([]string{\"IntValue=0\"}) == defaultStruct{IntValue: 0, intValue: 1, StringValue: \"a\", stringValue: \"a\"}, nil",
			args: args{
				arguments: []string{
					"IntValue=0",
				},
			},
			want: defaultStruct{
				IntValue:    0,
				intValue:    1,
				StringValue: "a",
				stringValue: "a",
			},
		},
		{
			name: "ArgumentParse([]string{\"intValue=0\"}) == defaultStruct{IntValue: 1, intValue: 0, StringValue: \"a\", stringValue: \"a\"}, nil",
			args: args{
				arguments: []string{
					"intValue=0",
				},
			},
			want: defaultStruct{
				IntValue:    1,
				intValue:    0,
				StringValue: "a",
				stringValue: "a",
			},
		},
		{
			name: "ArgumentParse([]string{\"StringValue=b\"}) == defaultStruct{IntValue: 1, intValue: 1, StringValue: \"b\", stringValue: \"a\"}, nil",
			args: args{
				arguments: []string{
					"StringValue=b",
				},
			},
			want: defaultStruct{
				IntValue:    1,
				intValue:    1,
				StringValue: "b",
				stringValue: "a",
			},
		},
		{
			name: "ArgumentParse([]string{\"stringValue=b\"}) == defaultStruct{IntValue: 1, intValue: 1, StringValue: \"a\", stringValue: \"b\"}, nil",
			args: args{
				arguments: []string{
					"stringValue=b",
				},
			},
			want: defaultStruct{
				IntValue:    1,
				intValue:    1,
				StringValue: "a",
				stringValue: "b",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ArgumentsParse[defaultStruct](tt.args.arguments)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArgumentParse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArgumentParse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArgumentsParse_invalidSliceType(t *testing.T) {
	type invalidSliceTypeStruct struct {
		SliceBoolValue []bool `key:"SliceBoolValue"`
	}
	type args struct {
		arguments []string
	}
	tests := []struct {
		name    string
		args    args
		want    invalidSliceTypeStruct
		wantErr bool
	}{
		{
			name: "ArgumentParse([]string{}) == invalidSliceTypeStruct{}, error",
			args: args{
				arguments: []string{
					"SliceBoolValue",
				},
			},
			want:    invalidSliceTypeStruct{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ArgumentsParse[invalidSliceTypeStruct](tt.args.arguments)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArgumentParse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArgumentParse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArgumentsParse_invalidType(t *testing.T) {
	type invalidTypeStruct struct {
		Float64Value float64 `key:"Float64Value"`
	}
	type args struct {
		arguments []string
	}
	tests := []struct {
		name    string
		args    args
		want    invalidTypeStruct
		wantErr bool
	}{
		{
			name: "ArgumentParse([]string{}) == invalidTypeStruct{}, error",
			args: args{
				arguments: []string{
					"Float64Value",
				},
			},
			want:    invalidTypeStruct{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ArgumentsParse[invalidTypeStruct](tt.args.arguments)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArgumentParse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArgumentParse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArgumentsParse_invalidDefaultValue(t *testing.T) {
	type invalidDefaultStruct struct {
		IntValue int `key:"IntValue" default:""`
	}
	type args struct {
		arguments []string
	}
	tests := []struct {
		name    string
		args    args
		want    invalidDefaultStruct
		wantErr bool
	}{
		{
			name: "ArgumentParse([]string{}) == invalidDefaultStruct{}, error",
			args: args{
				arguments: []string{},
			},
			want:    invalidDefaultStruct{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ArgumentsParse[invalidDefaultStruct](tt.args.arguments)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArgumentParse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArgumentParse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArgumentsParse_invalidDefaultType(t *testing.T) {
	type invalidDefaultTypeStruct struct {
		Float64Value float64 `key:"Float64Value" default:""`
	}
	type args struct {
		arguments []string
	}
	tests := []struct {
		name    string
		args    args
		want    invalidDefaultTypeStruct
		wantErr bool
	}{
		{
			name: "ArgumentParse([]string{}) == invalidDefaultTypeStruct{}, error",
			args: args{
				arguments: []string{},
			},
			want:    invalidDefaultTypeStruct{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ArgumentsParse[invalidDefaultTypeStruct](tt.args.arguments)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArgumentParse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArgumentParse() = %v, want %v", got, tt.want)
			}
		})
	}
}
