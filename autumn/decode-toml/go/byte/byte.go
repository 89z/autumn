package main

import (
   "decode/assert"
   "github.com/pelletier/go-toml"
   "io/ioutil"
   "log"
)

func main() {
   y, e := ioutil.ReadFile("Cargo.toml")
   if e != nil {
      log.Fatal(e)
   }
   m := assert.Map{}
   e = toml.Unmarshal(y, &m)
   if e != nil {
      log.Fatal(e)
   }
   s := m.M("package").S("edition")
   println(s == "2018")
}
