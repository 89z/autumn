package main

import (
   "encoding/json"
   "os"
)

func main() {
   s := "west & east"
   b, e := json.Marshal(s)
   if e != nil {
      panic(e)
   }
   os.Stdout.Write(append(b, '\n'))
}
