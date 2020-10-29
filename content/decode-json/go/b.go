package main

import (
   "encoding/json"
   "fmt"
)

func main() {
   y := []byte(`{"year": 2019, "month": 12}`)
   // example 1
   m := map[string]int{}
   json.Unmarshal(y, &m)
   fmt.Println(m)
   // example 2
   type Date struct {
      Year, Month int
   }
   o := Date{}
   json.Unmarshal(y, &o)
   fmt.Printf("%+v\n", o)
}

