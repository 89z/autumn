package main

import (
   "bufio"
   "os"
)

func main() {
   o_open, e := os.Open("a.go")
   if e != nil {
      os.Exit(1)
   }
   o_scan := bufio.NewScanner(o_open)
   for o_scan.Scan() {
      s := o_scan.Text()
      println(s)
   }
}
