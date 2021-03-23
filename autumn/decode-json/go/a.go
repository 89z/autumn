package main

import (
   "encoding/json"
   "fmt"
   "log"
)

func main() {
   b := []byte(`{"month": 12, "day": 31}`)
   var m struct { Month, Day int }
   e := json.Unmarshal(b, &m)
   if e != nil {
      log.Fatal(e)
   }
   fmt.Printf("%+v\n", m)
}
