package main

import (
   "decode/assert"
   "github.com/pelletier/go-toml"
   "io/ioutil"
   "log"
)

func TomlDecode(y []byte) (assert.Map, error) {
   o, e := toml.LoadBytes(y)
   if e != nil {
      return nil, e
   }
   return o.ToMap(), nil
}

func main() {
   y, e := ioutil.ReadFile("Cargo.toml")
   if e != nil {
      log.Fatal(e)
   }
   m, e := TomlDecode(y)
   if e != nil {
      log.Fatal(e)
   }
   s := m.A("package").M(0).S("name")
   println(s == "decode")
}
