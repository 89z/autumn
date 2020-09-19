package main

import (
   "bufio"
   "os"
)

func main() {
   o, e := os.Open("a.txt")
   if e != nil {
      os.Exit(1)
   }
   o1 := bufio.NewScanner(o)
   for o1.Scan() {
      s := o1.Text()
      println(s)
   }
}
