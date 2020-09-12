package main

import (
   "encoding/json"
   "fmt"
   "os"
)

func main() {
   m := map[string]int{"📗/📕": 10}
   y, e := json.Marshal(m)
   if e != nil {
      os.Exit(1)
   }
   fmt.Printf("%s\n", y)
}
