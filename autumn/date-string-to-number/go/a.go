package main

import (
   "os"
   "time"
)

func Parse(s string) (int64, error) {
   o, e := time.Parse(time.RFC3339[:len(s)], s)
   if e != nil {
      return 0, e
   }
   return o.Unix(), nil
}

func main() {
   n, e := Parse("2019-12-31")
   if e != nil {
      os.Exit(1)
   }
   println(n == 1577750400)
}
