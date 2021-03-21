package main

import (
   "bytes"
   "fmt"
)

func main() {
   b := []byte("May,June,July")
   a := bytes.SplitN(b, []byte{','}, 2)
   fmt.Printf("%q\n", a)
}
