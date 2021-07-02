package main

import (
   "bytes"
   "encoding/binary"
   "fmt"
)

func main() {
   var (
      b = new(bytes.Buffer)
      n uint8 = 0xFF
   )
   binary.Write(b, binary.BigEndian, n)
   fmt.Println(b.Bytes()) // [255]
}
