package main

import (
   "encoding/base64"
   "fmt"
)

func main() {
   s := "CgsMDQ=="
   a, e := base64.StdEncoding.DecodeString(s)
   if e != nil {
      panic(e)
   }
   fmt.Println(a)
}
