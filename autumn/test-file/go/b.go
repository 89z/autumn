package main

import (
   "log"
   "os"
)

func main() {
   o, e := os.Open("a.go")
   if os.IsNotExist(e) {
      log.Fatal(e)
   }
   o.Close()
}
