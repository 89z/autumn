package main

import (
   "fmt"
   "time"
)

func Format(o time.Duration) string {
   mil_n := o.Milliseconds() % 1000
   sec_n := int(o.Seconds()) % 60
   min_n := int(o.Minutes())
   return fmt.Sprintf("%v m %02v s %03v ms", min_n, sec_n, mil_n)
}

func main() {
   o := time.Now()
   for {
      time.Sleep(10 * time.Millisecond)
      s := Format(time.Since(o))
      fmt.Print("\r", s)
   }
}
