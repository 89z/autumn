package main

import (
   "encoding/json"
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
