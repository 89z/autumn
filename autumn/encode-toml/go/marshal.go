package main

import (
   "github.com/pelletier/go-toml"
   "log"
   "os"
)

type Map map[string]interface{}

func TomlEncode(source Map, dest string) (int, error) {
   y, e := toml.Marshal(source)
   if e != nil {
      return 0, e
   }
   o, e := os.Create(dest)
   if e != nil {
      return 0, e
   }
   return o.Write(y)
}

func main() {
   _, e := TomlEncode(Map{"month": 12, "day": 31}, "a.toml")
   if e != nil {
      log.Fatal(e)
   }
}
