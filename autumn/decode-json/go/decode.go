package main

import (
   "encoding/json"
   "strings"
)

func JsonDecode(s string) (Map, error) {
   o := strings.NewReader(s)
   m := Map{}
   e := json.NewDecoder(o).Decode(&m)
   if e != nil {
      return nil, e
   }
   return m, nil
}
