package main

import (
   "decode/assert"
   "github.com/pelletier/go-toml"
   "log"
   "os"
)

func main() {
   o, e := os.Open("Cargo.toml")
   if e != nil {
      log.Fatal(e)
   }
   m := assert.Map{}
   e = toml.NewDecoder(o).Decode(&m)
   if e != nil {
      log.Fatal(e)
   }
   s := m.M("package").S("edition")
   println(s == "2018")
}
