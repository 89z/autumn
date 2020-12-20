package main

import (
   "encoding/json"
   "log"
   "os"
)

func main() {
   y, e := json.Marshal(a)
   if e != nil {
      log.Fatal(e)
   }
   y = append(y, '\n')
   os.Stdout.Write(y)
}
