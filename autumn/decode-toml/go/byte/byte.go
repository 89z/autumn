package main

import (
   "decode/assert"
   "github.com/pelletier/go-toml"
   "io/ioutil"
   "log"
)

func TomlDecode(filename string) (assert.Map, error) {
   y, e := ioutil.ReadFile(filename)
   if e != nil {
      return nil, e
   }
   m := assert.Map{}
   return m, toml.Unmarshal(y, &m)
}

func main() {
   m, e := TomlDecode("Cargo.toml")
   if e != nil {
      log.Fatal(e)
   }
   s := m.M("package").S("edition")
   println(s == "2018")
}
