package main

import (
   "github.com/pelletier/go-toml"
   "os"
)

func TomlEncode(s string, m Map) error {
   o, e := os.Create(s)
   if e != nil {
      return e
   }
   return toml.NewEncoder(o).Encode(m)
}
