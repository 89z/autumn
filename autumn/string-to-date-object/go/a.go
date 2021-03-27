package main

import (
   "fmt"
   "time"
)

func date(s string) time.Time {
   var (
      year, day, hour, min, sec int
      mon time.Month
   )
   fmt.Sscanf(s, "%v-%v-%vT%v:%v:%v", &year, &mon, &day, &hour, &min, &sec)
   return time.Date(year, mon, day, hour, min, sec, 0, time.UTC)
}

func main() {
   for _, each := range []string{
      "2020-12-31T23:59:59",
      "2020-12-31T23:59",
      "2020-12-31T23",
      "2020-12-31",
   } {
      fmt.Println(date(each))
   }
}
