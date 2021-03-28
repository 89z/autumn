package main

import (
   "log"
   "os"
)

func main() {
   e := os.MkdirAll("north", os.ModeDir)
   if e != nil {
      log.Fatal(e)
   }
}
