package main

import (
   "fmt"
   "log"
   "os"
   "time"
)

func modTime(name string) (time.Time, error) {
   stat, e := os.Stat(name)
   if e != nil {
      return time.Time{}, e
   }
   return stat.ModTime(), nil
}

func main() {
   t, e := modTime("a.go")
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(t)
}
