package commandline

import (
	"fmt"
	"testing"
)

type argumentsStruct struct {
	intValue int `key:"intValue"`
}

type rootCommand struct{}

func (rootCommand) Execute(arguments []string) (string, error) {
	argumentsStruct, _ := ArgumentsParse[argumentsStruct](arguments)
	return fmt.Sprintf("rootCommand: {%d}", argumentsStruct.intValue), nil
}

type subCommand1 struct{}

func (subCommand1) execute(arguments []string) (string, error) {
	argumentsStruct, _ := ArgumentsParse[argumentsStruct](arguments)
	return fmt.Sprintf("subCommand1: {%d}", argumentsStruct.intValue), nil
}

func (s subCommand1) Execute(arguments []string) (string, error) {
	command, arguments := Find(s, arguments, subCommand2{})
	if command != s {
		return command.Execute(arguments)
	}
	return s.execute(arguments)
}
func (subCommand1) Name() string {
	return "subCommand1"
}

type subCommand2 struct{}

func (subCommand2) Execute(arguments []string) (string, error) {
	argumentsStruct, _ := ArgumentsParse[argumentsStruct](arguments)
	return fmt.Sprintf("subCommand2: {%d}", argumentsStruct.intValue), nil
}

func (subCommand2) Name() string {
	return "subCommand2"
}

func TestExecute(t *testing.T) {
	type args struct {
		rootCommand RootCommand
		arguments   []string
		subcommands []SubCommand
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Execute(rootCommand{}, nil) == \"rootCommand\": {0}",
			args: args{
				rootCommand: rootCommand{},
				arguments:   nil,
			},
			want:    "rootCommand: {0}",
			wantErr: false,
		},
		{
			name: "Execute(rootCommand{}, []string{\"intValue=1\"}) == \"rootCommand\": {1}",
			args: args{
				rootCommand: rootCommand{},
				arguments: []string{
					"intValue=1",
				},
			},
			want:    "rootCommand: {1}",
			wantErr: false,
		},
		{
			name: "Execute(rootCommand{}, []string{\"intValue=1\"}, []subCommand{subCommand1{}}) == \"rootCommand\": {1}",
			args: args{
				rootCommand: rootCommand{},
				arguments: []string{
					"intValue=1",
				},
				subcommands: []SubCommand{
					subCommand1{},
				},
			},
			want:    "rootCommand: {1}",
			wantErr: false,
		},
		{
			name: "Execute(rootCommand{}, []string{\"subCommand1\"}, []subCommand{subCommand1{}}) == \"subCommand1\": {0}",
			args: args{
				rootCommand: rootCommand{},
				arguments: []string{
					"subCommand1",
				},
				subcommands: []SubCommand{
					subCommand1{},
				},
			},
			want:    "subCommand1: {0}",
			wantErr: false,
		},
		{
			name: "Execute(rootCommand{}, []string{\"subCommand1\", \"intValue=1\"}, []subCommand{subCommand1{}}) == \"subCommand1\": {1}",
			args: args{
				rootCommand: rootCommand{},
				arguments: []string{
					"subCommand1",
					"intValue=1",
				},
				subcommands: []SubCommand{
					subCommand1{},
				},
			},
			want:    "subCommand1: {1}",
			wantErr: false,
		},
		{
			name: "Execute(rootCommand{}, []string{\"subCommand1\", \"subCommand2\"}, []subCommand{subCommand1{}}) == \"subCommand2\": {0}",
			args: args{
				rootCommand: rootCommand{},
				arguments: []string{
					"subCommand1",
					"subCommand2",
				},
				subcommands: []SubCommand{
					subCommand1{},
				},
			},
			want:    "subCommand2: {0}",
			wantErr: false,
		},
		{
			name: "Execute(rootCommand{}, []string{\"subCommand1\", \"subCommand2\", \"intValue=1\"}, []subCommand{subCommand1{}}) == \"subCommand2\": {1}",
			args: args{
				rootCommand: rootCommand{},
				arguments: []string{
					"subCommand1",
					"subCommand2",
					"intValue=1",
				},
				subcommands: []SubCommand{
					subCommand1{},
				},
			},
			want:    "subCommand2: {1}",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Execute(tt.args.rootCommand, tt.args.arguments, tt.args.subcommands...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
