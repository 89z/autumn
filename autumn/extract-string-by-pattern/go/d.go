package main

import (
   "bytes"
   "regexp"
)

func FindSubmatch(pat string, sub []byte) []byte {
   a := regexp.MustCompile(pat).FindSubmatch(sub)
   if len(a) < 2 {
      return []byte{}
   }
   return a[1]
}

func main() {
   y := FindSubmatch("a(..)", []byte("January"))
   println(bytes.Equal(y, []byte("nu")))
}
