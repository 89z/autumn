package main

import (
   "fmt"
   "time"
)

func main() {
   a := time.Now()
   for b := range time.Tick(100 * time.Millisecond) {
      fmt.Printf("%10s\r", b.Sub(a).Truncate(time.Millisecond))
   }
}
