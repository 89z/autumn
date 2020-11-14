package main

import (
   "log"
   "os"
)

func main() {
   in_o, e := os.Open("b.go")
   if e != nil {
      log.Fatal(e)
   }
   out_o, e := os.Create("c.go")
   if e != nil {
      log.Fatal(e)
   }
   out_o.ReadFrom(in_o)
}
