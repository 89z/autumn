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
   // example 1
   o = Date(2020, 5, 4, 3, 2, 1)
   fmt.Println(o)
   // example 2
   o = Date(2020, 5, 4, 3, 2)
   fmt.Println(o)
   // example 3
   o = Date(2020, 5, 4, 3)
   fmt.Println(o)
   // example 4
   o = Date(2020, 5, 4)
   fmt.Println(o)
   // example 5
   o = Date(2020, 5)
   fmt.Println(o)
   // example 6
   o = Date(2020)
   fmt.Println(o)
   // example 7
   o = Date()
   fmt.Println(o)
}
