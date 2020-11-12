package main

import (
   "bufio"
   "os"
)

func main() {
   open_o, e := os.Open("a.go")
   if e != nil {
      os.Exit(1)
   }
   scan_o := bufio.NewScanner(open_o)
   for scan_o.Scan() {
      s := scan_o.Text()
      println(s)
   }
}
