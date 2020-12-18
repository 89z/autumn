package main

import (
   "decode/assert"
   "github.com/pelletier/go-toml"
   "log"
   "strings"
)

func main() {
   o := strings.NewReader(`[package]
edition = "2018"
name = "decode-toml"
version = "1.0.0"`)
   m := assert.Map{}
   e := toml.NewDecoder(o).Decode(&m)
   if e != nil {
      log.Fatal(e)
   }
   s := m.M("package").S("edition")
   println(s == "2018")
}
