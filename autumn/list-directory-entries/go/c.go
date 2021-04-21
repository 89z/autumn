package main

import (
   "fmt"
   "os"
   "path/filepath"
)

func walk(root string) ([]string, error) {
   var a []string
   e := filepath.Walk(root, func(s string, f os.FileInfo, e error) error {
      if e != nil { return e }
      if ! f.IsDir() {
         a = append(a, s)
      }
      return nil
   })
   return a, e
}

func main() {
   a, e := walk("..")
   if e != nil {
      panic(e)
   }
   fmt.Println(a)
}
