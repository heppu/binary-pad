# Zero padding for integer binary presentations

I often find myself writing zero padding when debugging binary data, because viewing alligned binary strings in terminal is just easier.

Where as fmt.Printf("%b\%b\n%b\n%b\n", ^uint32(0), uint32(887678), uint32(1387), uint32(7945)) outputs:
```
11011000101101111110
10101101011
1111100001001
```

fmt.Printf("%b\%b\n%b\n%b\n",Pad(^uint32(0)), Pad(uint32(887678)), Pad(uint32(1387)), Pad(uint32(7945))) outputs:
```
11111111111111111111111111111111
00000000000011011000101101111110
00000000000000000000010101101011
00000000000000000001111100001001
```