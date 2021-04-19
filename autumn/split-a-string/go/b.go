package main

import (
   "bytes"
   "fmt"
)

func main() {
   b := []byte("west,north,east")
   a := bytes.SplitN(b, []byte{','}, 2)
   fmt.Printf("%q\n", a)
}
