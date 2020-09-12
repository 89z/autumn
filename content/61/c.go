package main

import (
   "os"
   "path/filepath"
)

func f(s string, o os.FileInfo, e error) error {
   if o.IsDir() {
      return e
   }
   println(s)
   return nil
}

func main() {
   filepath.Walk(".", f)
}
