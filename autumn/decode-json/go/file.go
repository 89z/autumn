package main

import (
   "encoding/json"
   "io/ioutil"
)

func JsonGetFile(s string) (Map, error) {
   y, e := ioutil.ReadFile(s)
   if e != nil {
      return nil, e
   }
   m := Map{}
   return m, json.Unmarshal(y, &m)
}
