package main

import (
   "fmt"
   "time"
)

func date(left ...int) time.Time {
   right := []int{1970, 1, 1, 0, 0, 0, 0}[len(left):]
   d := append(left, right...)
   return time.Date(
      d[0], time.Month(d[1]), d[2], d[3], d[4], d[5], d[6], time.UTC,
   )
}

func main() {
   var t time.Time
   // example 1
   t = date(2020, 5, 4, 3, 2, 1)
   fmt.Println(t)
   // example 2
   t = date(2020, 5, 4, 3, 2)
   fmt.Println(t)
   // example 3
   t = date(2020, 5, 4, 3)
   fmt.Println(t)
   // example 4
   t = date(2020, 5, 4)
   fmt.Println(t)
   // example 5
   t = date(2020, 5)
   fmt.Println(t)
   // example 6
   t = date(2020)
   fmt.Println(t)
   // example 7
   t = date()
   fmt.Println(t)
}
