package main

import (
   "encoding/hex"
   "fmt"
   "log"
)

func main() {
   s := "0a0b0c0d"
   a, e := hex.DecodeString(s)
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(a)
}
