package main

import (
   "os"
   "path/filepath"
)

func main() {
   sA := "index.md"
   sB, e := filepath.Abs(sA)
   if e != nil {
      os.Exit(1)
   }
   println(sB)
}
