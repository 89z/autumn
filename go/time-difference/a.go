package main

import (
   "fmt"
   "log"
   "time"
)

func f(s string) time.Time {
   n := len(s)
   s1 := time.RFC3339[:n]
   o, e := time.Parse(s1, s)
   if e != nil {
      log.Fatal(e)
   }
   return o
}

func main() {
   o := f("2019-12-31")
   o1 := f("2019-12-31T23:59:59")
   s := o1.Sub(o).String()
   fmt.Println(s == "23h59m59s")
}
