package main

import (
   "encoding/json"
   "net/http"
)

func JsonGetHttp(s string) (Map, error) {
   o, e := http.Get(s)
   if e != nil {
      return nil, e
   }
   m := Map{}
   return m, json.NewDecoder(o.Body).Decode(&m)
}
