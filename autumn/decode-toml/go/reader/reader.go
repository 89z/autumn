package main

import (
   "decode/assert"
   "github.com/pelletier/go-toml"
   "log"
   "os"
)

func TomlDecode(filename string) (assert.Map, error) {
   o, e := os.Open(filename)
   if e != nil {
      return nil, e
   }
   m := assert.Map{}
   return m, toml.NewDecoder(o).Decode(&m)
}

func main() {
   m, e := TomlDecode("Cargo.toml")
   if e != nil {
      log.Fatal(e)
   }
   s := m.M("package").S("edition")
   println(s == "2018")
}
