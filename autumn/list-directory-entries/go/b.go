package main

import (
   "fmt"
   "path/filepath"
)

func main() {
   a, e := filepath.Glob(`C:\*\LICENSE.txt`)
   if e != nil {
      panic(e)
   }
   fmt.Println(a)
}
