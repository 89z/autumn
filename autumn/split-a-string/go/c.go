package main

import (
   "encoding/csv"
   "fmt"
   "strings"
)

func main() {
   s := "north,east,south,west"
   r, e := csv.NewReader(strings.NewReader(s)).Read()
   if e != nil {
      panic(e)
   }
   fmt.Printf("%q\n", r)
}
