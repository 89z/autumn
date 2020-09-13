package main

import (
   "bufio"
   "fmt"
   "os"
)

func main() {
   o_open, e := os.Open("a.txt")
   if e != nil {
      os.Exit(1)
   }
   o_scan := bufio.NewScanner(o_open)
   a := []string{}
   for o_scan.Scan() {
      s := o_scan.Text()
      a = append(a, s)
   }
   fmt.Println(a)
}
