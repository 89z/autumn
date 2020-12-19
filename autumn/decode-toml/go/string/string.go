package main

import (
   "decode/assert"
   "github.com/pelletier/go-toml"
   "log"
   "strings"
)

func TomlDecode(s string) (assert.Map, error) {
   o := strings.NewReader(s)
   m := assert.Map{}
   return m, toml.NewDecoder(o).Decode(&m)
}

func main() {
   m, e := TomlDecode(`[package]
edition = "2018"
name = "decode-toml"
version = "1.0.0"`)
   if e != nil {
      log.Fatal(e)
   }
   s := m.M("package").S("edition")
   println(s == "2018")
}
