package main

import (
   "bytes"
   "regexp"
)

func FindSubmatch(str string, b []byte) []byte {
   a := regexp.MustCompile(str).FindSubmatch(b)
   if len(a) < 2 {
      return []byte{}
   }
   return a[1]
}

func main() {
   y := FindSubmatch("a(..)", []byte("January"))
   println(bytes.Equal(y, []byte("nu")))
}
