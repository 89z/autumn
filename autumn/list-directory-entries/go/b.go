package main

import (
   "fmt"
   "log"
   "os"
)

func main() {
   o, e := os.Open(".")
   if e != nil {
      log.Fatal(e)
   }
   a, e := o.Readdirnames(0)
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(a)
}
