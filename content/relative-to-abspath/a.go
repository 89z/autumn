package main

import (
   "os"
   "path/filepath"
)

func main() {
   s := "index.md"
   s2, e := filepath.Abs(s)
   if e != nil {
      os.Exit(1)
   }
   println(s2)
}
