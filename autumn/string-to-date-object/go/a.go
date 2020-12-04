package main

import (
   "fmt"
   "log"
   "time"
)

func Parse(value string) (time.Time, error) {
   n := len(value)
   layout := time.RFC3339[:n]
   return time.Parse(layout, value)
}

func main() {
   o, e := Parse("2019-12-31")
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(o)
}
