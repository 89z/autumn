package main

import (
   "encoding/json"
   "fmt"
   "log"
   "os"
)

func main() {
   // example 1
   s := `{"year": 2019, "month": 12}`
   y := []byte(s)
   a1 := map[string]int{}
   json.Unmarshal(y, &a1)
   // example 2
   o, e := os.Open("a.json")
   if e != nil {
      log.Fatal(e)
   }
   a2 := map[string]int{}
   json.NewDecoder(o).Decode(&a2)
   // print
   fmt.Print(a1, "\n", a2, "\n")
}
