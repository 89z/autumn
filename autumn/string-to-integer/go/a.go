package main

import (
   "fmt"
   "log"
)

func intVal(s string) (int, error) {
   var n int
   _, e := fmt.Sscan(s, &n)
   if e != nil {
      return 0, e
   }
   return n, nil
}

func main() {
   n, e := intVal("100")
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(n == 100)
}
