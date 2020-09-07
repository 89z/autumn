package main

import (
   "encoding/json"
   "fmt"
   "log"
   "os"
)

func main() {
   // string to map
   s := `{"year": 2019, "month": 12, "day": 31}`
   y := []byte(s)
   a := map[string]int{}
   json.Unmarshal(y, &a)
   // file to map
   o, e := os.Open("a.json")
   if e != nil {
      log.Fatal(e)
   }
   a2 := map[string]int{}
   json.NewDecoder(o).Decode(&a2)
   // print
   fmt.Print(a, "\n", a2, "\n")
}
