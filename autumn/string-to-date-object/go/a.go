package main

import (
   "fmt"
   "log"
   "time"
)

func ParseDate(value string) (time.Time, error) {
   n := len(value)
   layout := time.RFC3339[:n]
   return time.Parse(layout, value)
}

func main() {
   o, e := ParseDate("2019-12-31")
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(o)
}
