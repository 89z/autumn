package main

import (
   "bufio"
   "log"
   "os"
)

func main() {
   open, e := os.Open("a.go")
   if e != nil {
      log.Fatal(e)
   }
   scan := bufio.NewScanner(open)
   for scan.Scan() {
      s := scan.Text()
      println(s)
   }
}
