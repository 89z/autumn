package main

import (
   "decode/assert"
   "github.com/pelletier/go-toml"
   "log"
)

func TomlDecode(s string) (assert.Map, error) {
   o, e := toml.LoadFile(s)
   if e != nil {
      return nil, e
   }
   return o.ToMap(), nil
}

func main() {
   m, e := TomlDecode("Cargo.toml")
   if e != nil {
      log.Fatal(e)
   }
   s := m.A("package").M(0).S("name")
   println(s == "decode")
}
