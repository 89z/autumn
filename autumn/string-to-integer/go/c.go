package main

import (
   "log"
   "strconv"
)

func intVal(s string) (int64, error) {
   return strconv.ParseInt(s, 10, 64)
}

func main() {
   n, e := intVal("100")
   if e != nil {
      log.Fatal(e)
   }
   println(n == 100)
}
