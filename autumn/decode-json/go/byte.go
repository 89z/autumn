package main

import (
   "encoding/json"
   "io/ioutil"
   "log"
)

func JsonDecode(y []byte) (Map, error) {
   m := Map{}
   return m, json.Unmarshal(y, &m)
}

func main() {
   y, e := ioutil.ReadFile("a.json")
   if e != nil {
      log.Fatal(e)
   }
   m, e := JsonDecode(y)
   if e != nil {
      log.Fatal(e)
   }
   s := m.M("package").S("edition")
   println(s == "2019")
}
