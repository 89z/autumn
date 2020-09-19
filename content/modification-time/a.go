package main

import (
   "fmt"
   "os"
   "time"
)

func main() {
   s := "index.md"
   // example 1
   o1 := time.Now()
   os.Chtimes(s, o1, o1)
   // example 2
   o2, e := os.Stat(s)
   if e != nil {
      os.Exit(1)
   }
   // print
   n1 := o1.Unix()
   n2 := o2.ModTime().Unix()
   fmt.Println(n1 == n2)
}
