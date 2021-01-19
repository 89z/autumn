package main

import (
   "bytes"
   "regexp"
)

func main() {
   y := regexp.MustCompile("a..").Find([]byte("January"))
   println(bytes.Equal(y, []byte("anu")))
}
