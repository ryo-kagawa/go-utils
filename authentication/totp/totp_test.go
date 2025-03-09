package totp

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"testing"
)

func TestGenerate(t *testing.T) {
	sha1Key := []byte("12345678901234567890")
	sha256Key := []byte("12345678901234567890123456789012")
	sha512Key := []byte("1234567890123456789012345678901234567890123456789012345678901234")
	type args struct {
		key       []byte
		timeSteps int64
		digits    int
		crypto    func() hash.Hash
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Generate(sha1Key, 1, 8, sha1.New) == \"94287082\"",
			args: args{
				key:       sha1Key,
				timeSteps: 59 / 30,
				digits:    8,
				crypto:    sha1.New,
			},
			want: "94287082",
		},
		{
			name: "Generate(sha1Key, 37037036, 8, sha1.New) == \"07081804\"",
			args: args{
				key:       sha1Key,
				timeSteps: 1111111109 / 30,
				digits:    8,
				crypto:    sha1.New,
			},
			want: "07081804",
		},
		{
			name: "Generate(sha1Key, 37037037, 8, sha1.New) == \"14050471\"",
			args: args{
				key:       sha1Key,
				timeSteps: 1111111111 / 30,
				digits:    8,
				crypto:    sha1.New,
			},
			want: "14050471",
		},
		{
			name: "Generate(sha1Key, 41152263, 8, sha1.New) == \"89005924\"",
			args: args{
				key:       sha1Key,
				timeSteps: 1234567890 / 30,
				digits:    8,
				crypto:    sha1.New,
			},
			want: "89005924",
		},
		{
			name: "Generate(sha1Key, 66666666, 8, sha1.New) == \"69279037\"",
			args: args{
				key:       sha1Key,
				timeSteps: 2000000000 / 30,
				digits:    8,
				crypto:    sha1.New,
			},
			want: "69279037",
		},
		{
			name: "Generate(sha1Key, 666666666, 8, sha1.New) == \"65353130\"",
			args: args{
				key:       sha1Key,
				timeSteps: 20000000000 / 30,
				digits:    8,
				crypto:    sha1.New,
			},
			want: "65353130",
		},
		{
			name: "Generate(sha256Key, 1, 8, sha256.New) == \"46119246\"",
			args: args{
				key:       sha256Key,
				timeSteps: 59 / 30,
				digits:    8,
				crypto:    sha256.New,
			},
			want: "46119246",
		},
		{
			name: "Generate(sha256Key, 37037036, 8, sha256.New) == \"68084774\"",
			args: args{
				key:       sha256Key,
				timeSteps: 1111111109 / 30,
				digits:    8,
				crypto:    sha256.New,
			},
			want: "68084774",
		},
		{
			name: "Generate(sha256Key, 37037037, 8, sha256.New) == \"67062674\"",
			args: args{
				key:       sha256Key,
				timeSteps: 1111111111 / 30,
				digits:    8,
				crypto:    sha256.New,
			},
			want: "67062674",
		},
		{
			name: "Generate(sha256Key, 41152263, 8, sha256.New) == \"91819424\"",
			args: args{
				key:       sha256Key,
				timeSteps: 1234567890 / 30,
				digits:    8,
				crypto:    sha256.New,
			},
			want: "91819424",
		},
		{
			name: "Generate(sha256Key, 66666666, 8, sha256.New) == \"90698825\"",
			args: args{
				key:       sha256Key,
				timeSteps: 2000000000 / 30,
				digits:    8,
				crypto:    sha256.New,
			},
			want: "90698825",
		},
		{
			name: "Generate(sha256Key, 666666666, 8, sha256.New) == \"77737706\"",
			args: args{
				key:       sha256Key,
				timeSteps: 20000000000 / 30,
				digits:    8,
				crypto:    sha256.New,
			},
			want: "77737706",
		},
		{
			name: "Generate(sha512Key, 1, 8, sha512.New) == \"90693936\"",
			args: args{
				key:       sha512Key,
				timeSteps: 59 / 30,
				digits:    8,
				crypto:    sha512.New,
			},
			want: "90693936",
		},
		{
			name: "Generate(sha512Key, 37037036, 8, sha512.New) == \"25091201\"",
			args: args{
				key:       sha512Key,
				timeSteps: 1111111109 / 30,
				digits:    8,
				crypto:    sha512.New,
			},
			want: "25091201",
		},
		{
			name: "Generate(sha512Key, 37037037, 8, sha512.New) == \"99943326\"",
			args: args{
				key:       sha512Key,
				timeSteps: 1111111111 / 30,
				digits:    8,
				crypto:    sha512.New,
			},
			want: "99943326",
		},
		{
			name: "Generate(sha512Key, 41152263, 8, sha512.New) == \"93441116\"",
			args: args{
				key:       sha512Key,
				timeSteps: 1234567890 / 30,
				digits:    8,
				crypto:    sha512.New,
			},
			want: "93441116",
		},
		{
			name: "Generate(sha512Key, 66666666, 8, sha512.New) == \"38618901\"",
			args: args{
				key:       sha512Key,
				timeSteps: 2000000000 / 30,
				digits:    8,
				crypto:    sha512.New,
			},
			want: "38618901",
		},
		{
			name: "Generate(sha512Key, 666666666, 8, sha512.New) == \"47863826\"",
			args: args{
				key:       sha512Key,
				timeSteps: 20000000000 / 30,
				digits:    8,
				crypto:    sha512.New,
			},
			want: "47863826",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Generate(tt.args.key, tt.args.timeSteps, tt.args.digits, tt.args.crypto); got != tt.want {
				t.Errorf("Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
