package binary_test

import (
	"bytes"
	"testing"

	"github.com/ryo-kagawa/go-utils/binary"
)

func TestCounterBigEndian(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		value []byte
		want  []byte
	}{
		{
			name:  "CounterBigEndian(nil) == []byte{0x01}",
			value: nil,
			want:  []byte{0x01},
		},
		{
			name:  "CounterBigEndian([]byte(nil)) == []byte{0x01}",
			value: []byte(nil),
			want:  []byte{0x01},
		},
		{
			name:  "CounterBigEndian([]byte{0x00}]) == []byte{0x01}",
			value: []byte{0x00},
			want:  []byte{0x01},
		},
		{
			name:  "CounterBigEndian([]byte{0x01}]) == []byte{0x02}",
			value: []byte{0x01},
			want:  []byte{0x02},
		},
		{
			name:  "CounterBigEndian([]byte{0xff}]) == []byte{0x01, 0x00}",
			value: []byte{0xff},
			want:  []byte{0x01, 0x00},
		},
		{
			name:  "CounterBigEndian([]byte{0x00, 0x01}]) == []byte{0x00, 0x02}",
			value: []byte{0x00, 0x01},
			want:  []byte{0x00, 0x02},
		},
		{
			name:  "CounterBigEndian([]byte{0x01, 0x00}]) == []byte{0x01, 0x01}",
			value: []byte{0x01, 0x00},
			want:  []byte{0x01, 0x01},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binary.CounterBigEndian(tt.value); !bytes.Equal(got, tt.want) {
				t.Errorf("CounterBigEndian() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCounterLittleEndian(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		value []byte
		want  []byte
	}{
		{
			name:  "CounterLittleEndian(nil) == []byte{0x01}",
			value: nil,
			want:  []byte{0x01},
		},
		{
			name:  "CounterLittleEndian([]byte(nil)) == []byte{0x01}",
			value: []byte(nil),
			want:  []byte{0x01},
		},
		{
			name:  "CounterLittleEndian([]byte{0x00}]) == []byte{0x01}",
			value: []byte{0x00},
			want:  []byte{0x01},
		},
		{
			name:  "CounterLittleEndian([]byte{0x01}]) == []byte{0x02}",
			value: []byte{0x01},
			want:  []byte{0x02},
		},
		{
			name:  "CounterLittleEndian([]byte{0xff}]) == []byte{0x00, 0x01}",
			value: []byte{0xff},
			want:  []byte{0x00, 0x01},
		},
		{
			name:  "CounterLittleEndian([]byte{0x01, 0x00}]) == []byte{0x01, 0x00}",
			value: []byte{0x01, 0x00},
			want:  []byte{0x02, 0x00},
		},
		{
			name:  "CounterLittleEndian([]byte{0x00, 0x01}]) == []byte{0x01, 0x01}",
			value: []byte{0x00, 0x01},
			want:  []byte{0x01, 0x01},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binary.CounterLittleEndian(tt.value); !bytes.Equal(got, tt.want) {
				t.Errorf("CounterLittleEndian() = %v, want %v", got, tt.want)
			}
		})
	}
}
