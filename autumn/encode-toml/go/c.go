package main

import (
   "github.com/pelletier/go-toml"
   "os"
)

func TomlEncode(s string, m Map) error {
   y, e := toml.Marshal(m)
   if e != nil {
      return e
   }
   o, e := os.Create(s)
   if e != nil {
      return e
   }
   o.Write(y)
   return nil
}
