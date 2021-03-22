package main

import (
   "fmt"
   "log"
   "os"
   "path/filepath"
)

func walk(root string) ([]string, error) {
   var a []string
   e := filepath.Walk(root, func(s string, o os.FileInfo, e error) error {
      if e != nil { return e }
      if ! o.IsDir() {
         a = append(a, s)
      }
      return nil
   })
   return a, e
}

func main() {
   a, e := walk("..")
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(a)
}
