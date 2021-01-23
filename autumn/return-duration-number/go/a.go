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
   n, e := sinceHours("2020-12-31T01:02:31")
   check(n, e)
   n, e = sinceHours("2020-12-31T01:02")
   check(n, e)
   n, e = sinceHours("2020-12-31T01")
   check(n, e)
   n, e = sinceHours("2020-12-31")
   check(n, e)
   n, e = sinceHours("2020-12")
   check(n, e)
   n, e = sinceHours("2020")
   check(n, e)
}
