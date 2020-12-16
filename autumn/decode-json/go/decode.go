package main

import (
   "encoding/json"
   "strings"
)

func JsonDecode(s string) (Map, error) {
   o := strings.NewReader(s)
   m := Map{}
   return m, json.NewDecoder(o).Decode(&m)
}
