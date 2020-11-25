package main

import (
   "os"
   "path/filepath"
)

func main() {
   s, e := filepath.Abs("a.go")
   if e != nil {
      os.Exit(1)
   }
   println(s)
}
