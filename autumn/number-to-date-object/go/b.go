package main

import (
   "fmt"
   "time"
)

func Unix(sec int) time.Time {
   n := int64(sec)
   return time.Unix(n, 0)
}

func main() {
   n := 1609480799
   o := Unix(n)
   fmt.Println(o)
}
