package main

import (
   "github.com/pelletier/go-toml"
   "log"
   "os"
)

func main() {
   m := map[string]int{"month": 12, "day": 31}
   e := toml.NewEncoder(os.Stdout).Encode(m)
   if e != nil {
      log.Fatal(e)
   }
}
