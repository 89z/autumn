package main

import (
   "encoding/json"
   "fmt"
   "log"
   "os"
)

type Date struct {
   Year, Month int
}

func main() {
   // example 1
   s := `{"year": 2019, "month": 12}`
   y := []byte(s)
   a1 := Date{}
   json.Unmarshal(y, &a1)
   // example 2
   o, e := os.Open("a.json")
   if e != nil {
      log.Fatal(e)
   }
   a2 := Date{}
   json.NewDecoder(o).Decode(&a2)
   // print
   fmt.Printf("%+v\n%+v\n", a1, a2)
}
