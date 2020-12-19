package main

import (
   "github.com/pelletier/go-toml"
   "log"
   "os"
)

type Map map[string]interface{}

func TomlEncode(filename string, data Map) error {
   o, e := os.Create(filename)
   if e != nil {
      return e
   }
   return toml.NewEncoder(o).Encode(data)
}

func main() {
   m := Map{"month": 12, "day": 31}
   e := TomlEncode("a.toml", m)
   if e != nil {
      log.Fatal(e)
   }
}
