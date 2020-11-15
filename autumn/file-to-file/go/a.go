package main

import (
   "io"
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
   io.Copy(out_o, in_o)
}
