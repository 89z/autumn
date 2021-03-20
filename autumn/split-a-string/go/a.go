package main

import (
   "bytes"
   "fmt"
)

func main() {
   y := []byte("May,June,July")
   a := bytes.SplitN(y, []byte{','}, 2)
   fmt.Printf("%q\n", a)
}
