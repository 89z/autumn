package main

import (
   "io"
   "log"
   "os"
)

func main() {
   open, e := os.Open("a.go")
   if e != nil {
      log.Fatal(e)
   }
   data, e := io.ReadAll(open)
   if e != nil {
      log.Fatal(e)
   }
   os.Stdout.Write(data)
}
