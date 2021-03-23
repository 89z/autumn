package main

import (
   "encoding/base64"
   "fmt"
   "log"
)

func main() {
   s := "CgsMDQ=="
   a, e := base64.StdEncoding.DecodeString(s)
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(a)
}