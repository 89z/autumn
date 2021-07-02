package main

import (
   "encoding/binary"
   "fmt"
)

func main() {
   b := make([]byte, 2)
   binary.BigEndian.PutUint16(b, 0xFF)
   fmt.Println(b) // [0 255]
}
