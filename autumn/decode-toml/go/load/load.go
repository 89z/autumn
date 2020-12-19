package main

import (
   "decode/assert"
   "github.com/pelletier/go-toml"
   "log"
)

func TomlDecode(s string) (assert.Map, error) {
   o, e := toml.Load(s)
   if e != nil {
      return nil, e
   }
   return o.ToMap(), nil
}

func main() {
   m, e := TomlDecode(`[[package]]
name = "decode"
version = "1.0.0"`)
   if e != nil {
      log.Fatal(e)
   }
   s := m.A("package").M(0).S("name")
   println(s == "decode")
}
