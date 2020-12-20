package main

import (
   "fmt"
   "log"
   "os"
)

func main() {
   stat_o, e := os.Stat("a.go")
   if e != nil {
      log.Fatal(e)
   }
   mod_o := stat_o.ModTime()
   fmt.Println(mod_o)
}
