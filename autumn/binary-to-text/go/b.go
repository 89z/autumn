package main

import (
   "bytes"
   "encoding/ascii85"
)

func encode(src []byte) []byte {
   var dst bytes.Buffer
   ascii85.NewEncoder(&dst).Write(src)
   return dst.Bytes()
}

func main() {
   a := []byte{10, 11, 12, 13}
   b := encode(a)
   println(string(b) == "$4@7O")
}
