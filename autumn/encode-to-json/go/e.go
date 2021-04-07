package main

import (
   "encoding/json"
   "os"
)

func main() {
   a := []int{10, 11}
   b, e := json.MarshalIndent(a, "", "\t")
   if e != nil {
      panic(e)
   }
   os.Stdout.Write(append(b, '\n'))
}
