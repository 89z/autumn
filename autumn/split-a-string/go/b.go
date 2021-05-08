package main

import (
   "bytes"
   "fmt"
)

func main() {
   s, sep := []byte("west,north,east"), []byte{','}
   words := bytes.Split(s, sep)
   fmt.Printf("%c\n", words) // [[w e s t] [n o r t h] [e a s t]]
}
