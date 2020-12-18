package main

import (
   "encoding/json"
   "io"
   "log"
   "os"
)

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
