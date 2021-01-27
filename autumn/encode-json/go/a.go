package main

import (
   "encoding/json"
   "log"
   "os"
)

func main() {
   s := "May & June"
   e := json.NewEncoder(os.Stdout).Encode(s)
   if e != nil {
      log.Fatal(e)
   }
}
