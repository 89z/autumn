package main

import (
   "fmt"
   "path/filepath"
)

func main() {
   s, e := filepath.Glob(`C:\*\LICENSE.txt`)
   if e != nil {
      panic(e)
   }
   fmt.Println(s)
}
