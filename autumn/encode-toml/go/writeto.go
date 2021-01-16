package main

import (
   "github.com/pelletier/go-toml"
   "log"
   "os"
)

type Map map[string]interface{}

func TomlEncode(source Map, dest string) (int64, error) {
   tree, e := toml.TreeFromMap(source)
   if e != nil {
      return e
   }
   create, e := os.Create(dest)
   if e != nil {
      return e
   }
   return tree.WriteTo(create)
}

func main() {
   _, e := TomlEncode(Map{"month": 12, "day": 31}, "a.toml")
   if e != nil {
      log.Fatal(e)
   }
}
