package main

import (
   "fmt"
   "github.com/pelletier/go-toml"
   "log"
   "os"
)

type Map map[string]interface{}

func TomlEncode(s string, m Map) error {
   tree_o, e := toml.TreeFromMap(m)
   if e != nil {
      return e
   }
   create_o, e := os.Create(s)
   if e != nil {
      return e
   }
   n, e := tree_o.WriteTo(create_o)
   if e != nil {
      return fmt.Errorf("%v %v", n, e)
   }
   return nil
}

func main() {
   m := Map{"month": 12, "day": 31}
   e := TomlEncode("a.toml", m)
   if e != nil {
      log.Fatal(e)
   }
}
