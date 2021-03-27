package main

import (
   "fmt"
   "log"
   "time"
)

func parseDate(left string) (time.Time, error) {
   start := len(left)
   right := "1970-01-01T00:00:00Z"[start:]
   return time.Parse(time.RFC3339, left + right)
}

func main() {
   tests := []string{
      "",
      "2020",
      "2020-05",
      "2020-05-04",
      "2020-05-04T03",
      "2020-05-04T03:02",
      "2020-05-04T03:02:01",
   }
   for _, test := range tests {
      t, e := parseDate(test)
      if e != nil {
         log.Fatal(e)
      }
      fmt.Println(t)
   }
}
