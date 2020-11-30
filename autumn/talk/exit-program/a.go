package main

import (
   "log"
   "os"
)

func main() {
   o, e := os.Stat("a.go")
   if e != nil {
      log.Fatal(e)
   }
   n := o.Size()
   println(n)
}
