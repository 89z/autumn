package main

import (
   "github.com/pelletier/go-toml"
   "log"
)

type Map map[string]interface{}

func (m Map) M(s string) Map {
   return m[s].(map[string]interface{})
}

func (m Map) S(s string) string {
   return m[s].(string)
}

func main() {
   o, e := toml.LoadFile("Cargo.toml")
   if e != nil {
      log.Fatal(e)
   }
   m := Map{}
   o.Unmarshal(&m)
   s := m.M("package").S("edition")
   println(s == "2018")
}
