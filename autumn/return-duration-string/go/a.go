package main

import (
   "fmt"
   "time"
)

func format(d time.Duration) string {
   return fmt.Sprintf(
      "%v m %02v s %03v ms",
      int(d.Minutes()),
      int(d.Seconds()) % 60,
      d.Milliseconds() % 1000,
   )
}

func main() {
   t := time.Now()
   for {
      time.Sleep(10 * time.Millisecond)
      fmt.Print(
         "\r", format(time.Since(t)),
      )
   }
}
