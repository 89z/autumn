package main

import (
   "os"
   "time"
)

func ModTime(s string) (time.Time, error) {
   o, e := os.Stat(s)
   if e != nil {
      return time.Time{}, e
   }
   return o.ModTime(), nil
}

func main() {
   in_o := time.Now()
   os.Chtimes("a.go", in_o, in_o)
   out_o, e := ModTime("a.go")
   if e != nil {
      os.Exit(1)
   }
   println(in_o.Unix() == out_o.Unix())
}
