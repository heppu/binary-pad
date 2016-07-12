package pad

import (
	"fmt"
	"strings"
)

var bitCount int

func init() {
	x := ^uint(0)
	for ; x > 0; bitCount++ {
		x = x >> 1
	}
}

func Pad(i interface{}) (s string) {
	var bits int
	var isNegative bool

	switch i.(type) {
	case int:
		isNegative = i.(int) < 0
		bits = bitCount
	case uint:
		bits = bitCount
	case int8:
		isNegative = i.(int8) < 0
		bits = 8
	case uint8:
		bits = 8
	case int16:
		isNegative = i.(int16) < 0
		bits = 16
	case uint16:
		bits = 16
	case int32:
		isNegative = i.(int32) < 0
		bits = 32
	case uint32:
		bits = 32
	case int64:
		isNegative = i.(int64) < 0
		bits = 64
	case uint64:
		bits = 64
	default:
		return
	}

	s = fmt.Sprintf("%b", i)
	if isNegative {
		s = s[1:len(s)] // Remove the '-' prefix
		return strings.Repeat("1", bits-len(s)) + s
	}
	return strings.Repeat("0", bits-len(s)) + s
}
