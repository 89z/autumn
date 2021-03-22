package main

import (
   "fmt"
   "log"
   "os"
)

func read(name string) ([]string, error) {
   dir, e := os.Open(name)
   if e != nil { return nil, e }
   defer dir.Close()
   return dir.Readdirnames(0)
}

func main() {
   a, e := read(".")
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(a)
}
