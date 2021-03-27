package main

import (
   "encoding/json"
   "strings"
)

func main() {
   var (
      r = strings.NewReader(`{"month": 12, "day": 31}`)
      m struct { Month, Day int }
   )
   json.NewDecoder(r).Decode(&m)
   println(m.Day == 31)
}
