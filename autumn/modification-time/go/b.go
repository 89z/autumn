package main

import (
   "fmt"
   "os"
)

func main() {
   stat_o, e := os.Stat("a.go")
   if e != nil {
      os.Exit(1)
   }
   mod_o := stat_o.ModTime()
   fmt.Println(mod_o)
}
