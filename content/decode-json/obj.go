package main

import (
   "encoding/json"
   "fmt"
   "log"
   "os"
)

type Date struct {
   Year, Month, Day int
}

func main() {
   // string to object
   s := `{"year": 2019, "month": 12, "day": 31}`
   y := []byte(s)
   a := Date{}
   json.Unmarshal(y, &a)
   // file to object
   o, e := os.Open("a.json")
   if e != nil {
      log.Fatal(e)
   }
   a2 := Date{}
   json.NewDecoder(o).Decode(&a2)
   // print
   fmt.Printf("%+v\n%+v\n", a, a2)
}
