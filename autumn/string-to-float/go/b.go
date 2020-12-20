package main

import (
   "log"
   "strconv"
)

func FloatVal(s string) (float64, error) {
   return strconv.ParseFloat(s, 64)
}

func main() {
   n, e := FloatVal("9.9")
   if e != nil {
      log.Fatal(e)
   }
   println(n == 9.9)
}
