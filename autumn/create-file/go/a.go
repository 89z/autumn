package main

import (
   "log"
   "os"
)

func main() {
   f, e := os.OpenFile("b.go", os.O_CREATE | os.O_EXCL, os.ModePerm)
   if e != nil {
      log.Fatal(e)
   }
   f.Close()
}
