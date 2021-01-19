package main

import (
   "encoding/json"
   "io/ioutil"
)

func JsonLoadFile(path string) (Map, error) {
   data, e := ioutil.ReadFile(path)
   if e != nil {
      return nil, e
   }
   v := Map{}
   return v, json.Unmarshal(data, &v)
}
