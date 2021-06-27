package main

import (
   "io"
   "strings"
)

func scan(s io.Reader, sep byte) []byte {
   var (
      a [1]byte
      b []byte
   )
   for {
      _, err := s.Read(a[:])
      if err != nil {
         break
      } else if a[0] != sep {
         b = append(b, a[0])
      } else if b != nil {
         break
      }
   }
   return b
}

func main() {
   s := strings.NewReader(",north,,south,")
   for {
      text := scan(s, ',')
      if text == nil {
         break
      }
      println(string(text))
   }
}
