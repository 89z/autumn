package main

import (
   "fmt"
   "os"
)

func main() {
   o, e := os.Open("index.md")
   if e != nil {
      os.Exit(1)
   }
   y := make([]byte, 0x2000)
   for {
      n, e := o.Read(y)
      if e != nil {
         break
      }
      fmt.Printf("%s\n", y[:n])
   }
}
