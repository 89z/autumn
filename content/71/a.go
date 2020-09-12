package main

import (
   "io"
   "log"
   "os"
)

func main() {
   o_in, e := os.Open("a.txt")
   if e != nil {
      log.Fatal(e)
   }
   o_out, e := os.Create("b.txt")
   if e != nil {
      log.Fatal(e)
   }
   io.Copy(o_out, o_in)
}

