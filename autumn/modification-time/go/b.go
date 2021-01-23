package main

import (
   "fmt"
   "log"
   "os"
   "time"
)

func modTime(name string) (t time.Time, e error) {
   fi, e := os.Stat(name)
   if e != nil {
      return
   }
   return fi.ModTime(), nil
}

func main() {
   t, e := modTime("a.go")
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(t)
}
