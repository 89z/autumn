package main

import (
   "fmt"
   "log"
   "os"
)

func ReadDir(s string) ([]string, error) {
   o, e := os.Open(s)
   if e != nil {
      return nil, e
   }
   return o.Readdirnames(0)
}

func main() {
   a, e := ReadDir(".")
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(a)
}
