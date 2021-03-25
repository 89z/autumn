package main

import (
   "github.com/pelletier/go-toml"
   "log"
   "os"
)

func main() {
   m := map[string]int{"month": 12, "day": 31}
   b, e := toml.Marshal(m)
   if e != nil {
      log.Fatal(e)
   }
   os.Stdout.Write(b)
}
