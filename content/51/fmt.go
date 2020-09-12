package main

import (
   "fmt"
   "os"
)

func main() {
   s := ""
   n, e := fmt.Scanln(&s)
   if e != nil {
      os.Exit(1)
   }
   println(n, s)
}
