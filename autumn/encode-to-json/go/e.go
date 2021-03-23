package main

import (
   "encoding/json"
   "log"
   "os"
)

func main() {
   a := []int{10, 11}
   b, e := json.MarshalIndent(a, "", "\t")
   if e != nil {
      log.Fatal(e)
   }
   os.Stdout.Write(append(b, '\n'))
}
