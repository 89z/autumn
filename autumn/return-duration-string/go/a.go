package main

import (
   "fmt"
   "time"
)

func format(d time.Duration) string {
   m, s, ms := d.Minutes(), d.Seconds(), d.Milliseconds()
   return fmt.Sprintf(
      "%v m %02v s %03v ms", int(m), int(s) % 60, ms % 1000,
   )
}

func main() {
   t := time.Now()
   for {
      time.Sleep(10 * time.Millisecond)
      s := time.Since(t)
      fmt.Print(
         "\r", format(s),
      )
   }
}
