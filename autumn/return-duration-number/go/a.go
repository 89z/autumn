package main

import (
   "fmt"
   "log"
)

func check(n float64, e error) {
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(n)
}

func main() {
   n, e := TimeHours("2020-12-31T23:59:59")
   check(n, e)
   n, e = TimeHours("2020-12-31T23:59")
   check(n, e)
   n, e = TimeHours("2020-12-31T23")
   check(n, e)
   n, e = TimeHours("2020-12-31")
   check(n, e)
   n, e = TimeHours("2020-12")
   check(n, e)
   n, e = TimeHours("2020")
   check(n, e)
   n, e = TimeHours("")
   check(n, e)
}
