package main

import (
   "log"
   "os"
)

func main() {
   e := os.Link("a.go", "b.go")
   if e != nil {
      log.Fatal(e)
   }
}
