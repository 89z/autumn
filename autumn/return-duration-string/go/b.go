package main

import (
   "fmt"
   "time"
)

func main() {
   old_o := time.Now()
   for new_o := range time.Tick(10 * time.Millisecond) {
      n := new_o.Sub(old_o).Seconds()
      fmt.Printf("\r%.2f", n)
   }
}
