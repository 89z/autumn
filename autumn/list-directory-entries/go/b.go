package main

import (
   "log"
   "os"
)

func main() {
   dir, e := os.ReadDir(".")
   if e != nil {
      log.Fatal(e)
   }
   for _, each := range dir {
      println(each.Name())
   }
}
