package main

import (
   "log"
   "os"
)

func fileSize(name string) (int64, error) {
   info, e := os.Stat(name)
   if e != nil {
      return 0, e
   }
   return info.Size(), nil
}

func main() {
   n, e := fileSize("a.go")
   if e != nil {
      log.Fatal(e)
   }
   println(n)
}
