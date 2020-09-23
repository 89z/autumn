package main

import (
   "log"
   "os"
   "path/filepath"
)

func f(s string, o os.FileInfo, e error) error {
   if e != nil {
      return e
   }
   if ! o.IsDir() {
      println(s)
   }
   return nil
}

func main() {
   e := filepath.Walk(".", f)
   if e != nil {
      log.Fatal(e)
   }
}
