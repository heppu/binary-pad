package pad

import (
	"testing"
)

var int8Values = []struct {
	in  int8
	out string
}{
	{-128, "10000000"},
	{127, "01111111"},
	{0, "00000000"},
	{1, "00000001"},
	{-1, "11111111"},
}

var int16Values = []struct {
	in  int16
	out string
}{
	{-32768, "1000000000000000"},
	{32767, "0111111111111111"},
	{0, "0000000000000000"},
	{1, "0000000000000001"},
	{-1, "1111111111111111"},
}

var uint32Values = []struct {
	in  uint32
	out string
}{
	{4294967295, "11111111111111111111111111111111"},
	{0, "00000000000000000000000000000000"},
	{1, "00000000000000000000000000000001"},
}

func TestPadInt8(t *testing.T) {
	var s string
	for _, tt := range int8Values {
		s = Pad(tt.in)
		if s != tt.out {
			t.Errorf("Failed with: %d. Expected: %s got: %s", tt.in, tt.out, s)
		}
	}
}

func TestPadInt16(t *testing.T) {
	var s string
	for _, tt := range int16Values {
		s = Pad(tt.in)
		if s != tt.out {
			t.Errorf("Failed with: %d. Expected: %s got: %s", tt.in, tt.out, s)
		}
	}
}

func TestPadUint32(t *testing.T) {
	var s string
	for _, tt := range uint32Values {
		s = Pad(tt.in)
		if s != tt.out {
			t.Errorf("Failed with: %d. Expected: %s got: %s", tt.in, tt.out, s)
		}
	}
}
