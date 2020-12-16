package main

import (
   "fmt"
   "log"
)

func main() {
   n, e := Hours("2019-12-31")
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(n)
}
