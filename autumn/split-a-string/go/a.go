package main

import (
   "bytes"
   "fmt"
)

func main() {
   y := []byte("May,June,July")
   a := bytes.Split(y, []byte{','})
   fmt.Printf("%q\n", a)
}
