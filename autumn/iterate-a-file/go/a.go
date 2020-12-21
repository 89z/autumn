package main

import (
   "bufio"
   "log"
   "os"
)

func main() {
   open_o, e := os.Open("a.go")
   if e != nil {
      log.Fatal(e)
   }
   scan_o := bufio.NewScanner(open_o)
   for scan_o.Scan() {
      s := scan_o.Text()
      println(s)
   }
}
