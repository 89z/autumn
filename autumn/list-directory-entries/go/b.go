package main

import (
   "fmt"
   "log"
   "os"
)

func ReadDir(s string) ([]string, error) {
   dir, e := os.Open(s)
   if e != nil {
      return nil, e
   }
   defer dir.Close()
   return dir.Readdirnames(0)
}

func main() {
   a, e := ReadDir(".")
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(a)
}
