package main

import (
   "fmt"
   "time"
)

func Date(left ...int) time.Time {
   right := []int{1970, 1, 1, 0, 0, 0, 0}[len(left):]
   d := append(left, right...)
   return time.Date(
      d[0], time.Month(d[1]), d[2], d[3], d[4], d[5], d[6], time.UTC,
   )
}

func main() {
   var o time.Time
   o = Date(2020, 12, 31, 23, 59, 59)
   fmt.Println(o)
   o = Date(2020, 12, 31, 23, 59)
   fmt.Println(o)
   o = Date(2020, 12, 31, 23)
   fmt.Println(o)
   o = Date(2020, 12, 31)
   fmt.Println(o)
   o = Date(2020, 12)
   fmt.Println(o)
   o = Date(2020)
   fmt.Println(o)
   o = Date()
   fmt.Println(o)
}
