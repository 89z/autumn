package main

import (
   "io"
   "log"
   "os"
)

func main() {
   in_o, e := os.Open("a.txt")
   if e != nil {
      log.Fatal(e)
   }
   out_o, e := os.Create("b.txt")
   if e != nil {
      log.Fatal(e)
   }
   io.Copy(out_o, in_o)
}

