package main

import (
   "log"
   "os"
)

func fileSize(name string) (int64, error) {
   fi, err := os.Stat(name)
   if err != nil {
      return 0, err
   }
   return fi.Size(), nil
}

func main() {
   n, e := fileSize("a.go")
   if e != nil {
      log.Fatal(e)
   }
   println(n)
}
