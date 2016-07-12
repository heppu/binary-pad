package main

import (
	"fmt"
	"github.com/heppu/binary-pad"
	"math/rand"
	"time"
)

const BUF_LEN = 5

func main() {
	rand.Seed(time.Now().UnixNano())
	buf := make([]byte, BUF_LEN)
	readBytesFromDevice(&buf)

	// Print buffer zero padded
	for _, v := range buf {
		fmt.Println(pad.Pad(v))
	}
}

// Fake reading data from device
func readBytesFromDevice(b *[]byte) {
	for i := 0; i < BUF_LEN; i++ {
		(*b)[i] = byte(rand.Intn(int(^uint8(0))))
	}
}
