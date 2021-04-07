package main

import (
   "encoding/hex"
   "fmt"
)

func main() {
   s := "0a0b0c0d"
   a, e := hex.DecodeString(s)
   if e != nil {
      panic(e)
   }
   fmt.Println(a)
}
