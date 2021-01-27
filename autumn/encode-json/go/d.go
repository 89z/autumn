package main

import (
   "encoding/json"
   "log"
   "os"
)

func main() {
   s := "May & June"
   b, e := json.Marshal(s)
   if e != nil {
      log.Fatal(e)
   }
   os.Stdout.Write(append(b, '\n'))
}
