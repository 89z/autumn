package main

import (
   "encoding/json"
   "log"
   "os"
)

func JsonDecode(s string) (Map, error) {
   o, e := os.Open(s)
   if e != nil {
      return nil, e
   }
   m := Map{}
   e = json.NewDecoder(o).Decode(&m)
   if e != nil {
      return nil, e
   }
   return m, nil
}

func main() {
   m, e := JsonDecode("youtube.json")
   if e != nil {
      log.Fatal(e)
   }
   s := m.M("videoDetails").A("keywords")[0]
   println(s == "Blue Wednesday")
}
