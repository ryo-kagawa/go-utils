package strings

import (
	"testing"
)

func TestPad(t *testing.T) {
	type args struct {
		str       string
		padString string
		length    int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "func(\"a\", \"b\", 1) == \"\"",
			args: args{
				str:       "a",
				padString: "b",
				length:    1,
			},
			want: "",
		},
		{
			name: "func(\"a\", \"b\", 2) == \"\"",
			args: args{
				str:       "a",
				padString: "b",
				length:    2,
			},
			want: "b",
		},
		{
			name: "func(\"a\", \"b\", 2) == \"\"",
			args: args{
				str:       "a",
				padString: "bc",
				length:    4,
			},
			want: "bcb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pad(tt.args.str, tt.args.padString, tt.args.length); got != tt.want {
				t.Errorf("pad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPadLeft(t *testing.T) {
	type args struct {
		str       string
		padString string
		length    int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "func(\"aaa\", \"b\", 1) == \"aaa\"",
			args: args{
				str:       "aaa",
				padString: "b",
				length:    1,
			},
			want: "aaa",
		},
		{
			name: "func(\"aaa\", \"b\", 6) == \"bbbaaa\"",
			args: args{
				str:       "aaa",
				padString: "b",
				length:    6,
			},
			want: "bbbaaa",
		},
		{
			name: "func(\"aaa\", \"b\", 6) == \"bcbaaa\"",
			args: args{
				str:       "aaa",
				padString: "bc",
				length:    6,
			},
			want: "bcbaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PadLeft(tt.args.str, tt.args.padString, tt.args.length); got != tt.want {
				t.Errorf("PadLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPadRight(t *testing.T) {
	type args struct {
		str       string
		padString string
		length    int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "func(\"aaa\", \"b\", 1) == \"aaa\"",
			args: args{
				str:       "aaa",
				padString: "b",
				length:    1,
			},
			want: "aaa",
		},
		{
			name: "func(\"aaa\", \"b\", 6) == \"aaabbb\"",
			args: args{
				str:       "aaa",
				padString: "b",
				length:    6,
			},
			want: "aaabbb",
		},
		{
			name: "func(\"aaa\", \"b\", 6) == \"aaabcb\"",
			args: args{
				str:       "aaa",
				padString: "bc",
				length:    6,
			},
			want: "aaabcb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PadRight(tt.args.str, tt.args.padString, tt.args.length); got != tt.want {
				t.Errorf("PadRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
