package main

import (
   "fmt"
   "time"
)

func main() {
   o := time.Now()
   for {
      time.Sleep(10 * time.Millisecond)
      n := time.Since(o).Seconds()
      fmt.Printf("\r%.2f", n)
   }
}
