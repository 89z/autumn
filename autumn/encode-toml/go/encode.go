package main

import (
   "github.com/pelletier/go-toml"
   "log"
   "os"
)

type Map map[string]interface{}

func TomlEncode(source Map, dest string) error {
   o, e := os.Create(dest)
   if e != nil {
      return e
   }
   return toml.NewEncoder(o).Encode(source)
}

func main() {
   e := TomlEncode(Map{"month": 12, "day": 31}, "a.toml")
   if e != nil {
      log.Fatal(e)
   }
}
