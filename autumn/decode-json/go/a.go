package main

import (
   "encoding/json"
   "io"
   "log"
   "os"
)

type Map map[string]interface{}

func (m Map) M(s string) Map {
   return m[s].(map[string]interface{})
}

func (m Map) S(s string) string {
   return m[s].(string)
}

func JsonDecode(o io.Reader) (Map, error) {
   m := Map{}
   return m, json.NewDecoder(o).Decode(&m)
}

func main() {
   o, e := os.Open("a.json")
   if e != nil {
      log.Fatal(e)
   }
   m, e := JsonDecode(o)
   if e != nil {
      log.Fatal(e)
   }
   s := m.M("package").S("edition")
   println(s == "2019")
}
