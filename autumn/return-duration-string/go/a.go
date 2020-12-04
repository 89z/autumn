package main

import (
   "fmt"
   "time"
)

func main() {
   o := time.Now()
   for {
      time.Sleep(100 * time.Millisecond)
      s := time.Now().Sub(o).Truncate(10 * time.Millisecond)
      fmt.Printf("%10s\r", s)
   }
}
