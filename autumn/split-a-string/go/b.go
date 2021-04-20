package main

import (
   "bytes"
   "fmt"
)

func main() {
   s, sep := []byte("west,north,east"), []byte{','}
   sub := bytes.SplitN(s, sep, 2)
   fmt.Printf("%q\n", sub) // ["west" "north,east"]
}
