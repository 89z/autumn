package main

import (
   "log"
   "os"
)

func main() {
   e := os.MkdirAll("March", os.ModeDir)
   if e != nil {
      log.Fatal(e)
   }
}
