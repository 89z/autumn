package main

import (
   "fmt"
   "log"
)

func main() {
   n, e := FloatVal("9.9")
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(n == 9.9)
}
