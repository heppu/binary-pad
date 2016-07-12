# Zero padding for integer binary presentations

I often find myself writing zero padding when debugging binary data, because viewing alligned binary strings in terminal is just easier.

##Comparison

####Without Pad:
```go
fmt.Printf("%b\n%b\n%b\n%b\n", ^uint32(0), uint32(887678), uint32(1387), uint32(7945))
11011000101101111110
10101101011
1111100001001
```

####With Pad:
```go
fmt.Printf("%s\n%s\n%s\n%s\n",Pad(^uint32(0)), Pad(uint32(887678)), Pad(uint32(1387)), Pad(uint32(7945)))
11111111111111111111111111111111
00000000000011011000101101111110
00000000000000000000010101101011
00000000000000000001111100001001
```

##Usage

```go
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
```
