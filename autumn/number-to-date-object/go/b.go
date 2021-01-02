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
   n := 1577858399
   o := Unix(n)
   fmt.Println(o)
}
