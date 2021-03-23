package main

import (
   "encoding/json"
   "fmt"
   "log"
   "strings"
)

func main() {
   s := strings.NewReader(`{"month": 12, "day": 31}`)
   var m struct { Month, Day int }
   e := json.NewDecoder(s).Decode(&m)
   if e != nil {
      log.Fatal(e)
   }
   fmt.Printf("%+v\n", m)
}
