package main

import (
   "fmt"
   "log"
   "time"
)

func ParseDate(value string) (time.Time, error) {
   layout := time.RFC3339[:len(value)]
   return time.Parse(layout, value)
}

func main() {
   o, e := ParseDate("2020-12-31")
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(o)
}
