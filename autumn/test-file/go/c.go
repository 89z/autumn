package main

import (
   "log"
   "os"
)

func main() {
   o, e := os.OpenFile("d.go", os.O_CREATE | os.O_EXCL, os.ModePerm)
   if os.IsExist(e) {
      log.Fatal(e)
   }
   o.Close()
}
