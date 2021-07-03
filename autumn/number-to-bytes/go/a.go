package main

import (
   "encoding/binary"
   "fmt"
   "math/bits"
)

func encodeUint(x uint64) []byte {
   buf := make([]byte, 8)
   binary.BigEndian.PutUint64(buf, x)
   return buf[bits.LeadingZeros64(x) >> 3:]
}

func main() {
   buf := encodeUint(0xFF)
   fmt.Println(buf) // [255]
}
