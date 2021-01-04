package main

import (
   "github.com/pelletier/go-toml"
   "os"
)

func TomlEncode(s string, m Map) error {
   tree_o, e := toml.TreeFromMap(m)
   if e != nil {
      return e
   }
   create_o, e := os.Create(s)
   if e != nil {
      return e
   }
   tree_o.WriteTo(create_o)
   return nil
}
