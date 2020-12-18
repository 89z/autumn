package main

import (
   "encoding/json"
   "log"
)

func JsonDecode(y []byte) (Map, error) {
   m := Map{}
   return m, json.Unmarshal(y, &m)
}

func main() {
   y := []byte(`{"package": {"edition": "2019"}}`)
   m, e := JsonDecode(y)
   if e != nil {
      log.Fatal(e)
   }
   s := m.M("package").S("edition")
   println(s == "2019")
}
