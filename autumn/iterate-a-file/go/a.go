package main

import (
   "bufio"
   "log"
   "os"
)

func main() {
   file, e := os.Open("a.go")
   if e != nil {
      log.Fatal(e)
   }
   defer file.Close()
   buf := bufio.NewScanner(file)
   for buf.Scan() {
      println(buf.Text())
   }
}
