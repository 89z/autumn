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
   n, e := TimeHours("2020-05-04T03:02:01")
   check(n, e)
   n, e = TimeHours("2020-05-04T03:02")
   check(n, e)
   n, e = TimeHours("2020-05-04T03")
   check(n, e)
   n, e = TimeHours("2020-05-04")
   check(n, e)
   n, e = TimeHours("2020-05")
   check(n, e)
   n, e = TimeHours("2020")
   check(n, e)
   n, e = TimeHours("")
   check(n, e)
}
