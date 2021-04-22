package main

import (
   "io/fs"
   "path/filepath"
)

func walk(s string, d fs.DirEntry, e error) error {
   if e != nil { return e }
   if ! d.IsDir() {
      println(s)
   }
   return nil
}

func main() {
   filepath.WalkDir("..", walk)
}
