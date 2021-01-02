package main

import (
   "fmt"
   "log"
   "time"
)

func check(o time.Time, e error) {
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(o)
}

func ParseDate(left string) (time.Time, error) {
   start := len(left)
   right := "1970-01-01T00:00:00Z"[start:]
   return time.Parse(time.RFC3339, left + right)
}

func main() {
   o, e := ParseDate("2020-12-31T23:59:59")
   check(o, e)
   o, e = ParseDate("2020-12-31T23:59")
   check(o, e)
   o, e = ParseDate("2020-12-31T23")
   check(o, e)
   o, e = ParseDate("2020-12-31")
   check(o, e)
   o, e = ParseDate("2020-12")
   check(o, e)
   o, e = ParseDate("2020")
   check(o, e)
   o, e = ParseDate("")
   check(o, e)
}
