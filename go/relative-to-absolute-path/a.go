package main

import (
   "os"
   "path/filepath"
)

func main() {
   s := "index.md"
   s1, e := filepath.Abs(s)
   if e != nil {
      os.Exit(1)
   }
   println(s1)
}
