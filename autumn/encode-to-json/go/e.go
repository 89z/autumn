package main

import (
   "encoding/json"
   "os"
)

func main() {
   s := []int{10, 11}
   b, e := json.MarshalIndent(s, "", " ")
   if e != nil {
      panic(e)
   }
   os.Stdout.Write(append(b, '\n'))
}
