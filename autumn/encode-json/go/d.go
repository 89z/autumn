package main

import (
   "encoding/json"
   "fmt"
   "log"
)

func main() {
   y, e := json.Marshal(a)
   if e != nil {
      log.Fatal(e)
   }
   fmt.Printf("%s\n", y)
}
