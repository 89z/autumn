package main

import (
   "encoding/csv"
   "fmt"
   "strings"
)

func main() {
   {
      r := csv.NewReader(strings.NewReader("west,north,east"))
      s, e := r.Read()
      if e != nil {
         panic(e)
      }
      fmt.Println(s)
   }
   {
      r := csv.NewReader(strings.NewReader("west north east"))
      r.Comma = ' '
      s, e := r.Read()
      if e != nil {
         panic(e)
      }
      fmt.Println(s)
   }
}
