package main

import (
   "bytes"
   "regexp"
)

func findSubmatch(re string, input []byte) []byte {
   a := regexp.MustCompile(re).FindSubmatch(input)
   if len(a) < 2 {
      return []byte{}
   }
   return a[1]
}

func main() {
   y := findSubmatch("a(..)", []byte("January"))
   println(bytes.Equal(y, []byte("nu")))
}
