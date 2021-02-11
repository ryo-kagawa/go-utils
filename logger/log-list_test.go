package logger

import (
	"strings"
	"testing"
)

type testStruct struct {
	s  string
	i  int
	sa []string
}

func Test_logList_generateLog(t *testing.T) {
	logList := &logList{}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "string",
			args: args{
				v: "s",
			},
			want:    `"s"`,
			wantErr: false,
		},
		{
			name: "int",
			args: args{
				v: 1,
			},
			want:    `1`,
			wantErr: false,
		},
		{
			name: "stringArray",
			args: args{
				v: []string{
					"s1",
					"S2",
				},
			},
			want:    `["s1","s2"]`,
			wantErr: false,
		},
		{
			name: "int",
			args: args{
				v: 1,
			},
			want:    `1`,
			wantErr: false,
		},
		{
			name: "struct",
			args: args{
				v: testStruct{
					s: "string",
					i: 1,
					sa: []string{
						"sa1",
						"sa2",
					},
				},
			},
			want:    `{"s":"string","i":1,"sa":["sa1","sa2"]}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := logList.generateLog(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("logList.generateLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if strings.HasPrefix(got, `"message":`+tt.want+`}`) {
				t.Errorf("logList.generateLog() = strings.HasPrefix(%v, %v)", got, tt.want)
			}
		})
	}
}
