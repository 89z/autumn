package main

import (
   "github.com/pelletier/go-toml"
   "log"
   "os"
)

func main() {
   m := map[string]interface{}{"month": 12, "day": 31}
   o, e := toml.TreeFromMap(m)
   if e != nil {
      log.Fatal(e)
   }
   n, e := o.WriteTo(os.Stdout)
   if e != nil {
      log.Fatal(n, e)
   }
}
