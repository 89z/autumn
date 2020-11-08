package main

import (
   "encoding/json"
   "fmt"
   "strings"
)

func main() {
   s := `{"year": 2019, "month": 12}`
   // example 1
   m := map[string]int{}
   json.NewDecoder(strings.NewReader(s)).Decode(&m)
   fmt.Println(m)
   // example 2
   type Date struct {
      Year, Month int
   }
   o := Date{}
   json.NewDecoder(strings.NewReader(s)).Decode(&o)
   fmt.Printf("%+v\n", o)
}
