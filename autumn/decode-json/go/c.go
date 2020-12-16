package main

import (
   "encoding/json"
   "io/ioutil"
)

func JsonDecode(s string) (Map, error) {
   y, e := ioutil.ReadFile(s)
   if e != nil {
      return nil, e
   }
   m := Map{}
   e = json.Unmarshal(y, &m)
   if e != nil {
      return nil, e
   }
   return m, nil
}
