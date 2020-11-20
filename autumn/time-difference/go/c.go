package main

import (
   "fmt"
   "math"
   "time"
)

func main() {
   old_o := time.Now()
   for new_o := range time.Tick(10 * time.Millisecond) {
      diff_n := new_o.Sub(old_o).Seconds()
      min_n := int(diff_n) / 60
      sec_n := math.Mod(diff_n, 60)
      fmt.Printf("%d m %5.2f s\r", min_n, sec_n)
   }
}
