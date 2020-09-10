package main

import (
   "fmt"
   "os"
   "time"
)

func main() {
   o := time.Now()
   os.Chtimes("index.md", o, o)
   o2, e := os.Stat("index.md")
   if e != nil {
      os.Exit(1)
   }
   n := o.Unix()
   n2 := o2.ModTime().Unix()
   fmt.Println(n, n2 == n)
}
