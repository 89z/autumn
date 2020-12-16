package main

import (
   "fmt"
   "log"
)

func IntVal(s string) (int, error) {
   var value int
   count, e := fmt.Sscan(s, &value)
   if e != nil {
      return 0, fmt.Errorf("%v %v", count, e)
   }
   return value, nil
}

func main() {
   n, e := IntVal("100")
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(n == 100)
}
