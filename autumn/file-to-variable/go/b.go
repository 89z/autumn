package main

import (
   "log"
   "os"
)

func main() {
   data, e := os.ReadFile("a.go")
   if e != nil {
      log.Fatal(e)
   }
   os.Stdout.Write(data)
}
