package main

import (
   "encoding/json"
   "log"
   "strings"
)

func JsonDecode(s string) (Map, error) {
   o := strings.NewReader(s)
   m := Map{}
   return m, json.NewDecoder(o).Decode(&m)
}

func main() {
   m, e := JsonDecode(`{"package": {"edition": "2019"}}`)
   if e != nil {
      log.Fatal(e)
   }
   s := m.M("package").S("edition")
   println(s == "2019")
}
