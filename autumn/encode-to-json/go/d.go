package main

import (
   "encoding/json"
   "log"
   "os"
)

func main() {
   s := "west & east"
   b, e := json.Marshal(s)
   if e != nil {
      log.Fatal(e)
   }
   os.Stdout.Write(append(b, '\n'))
}
