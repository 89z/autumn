package main

import (
   "encoding/csv"
   "fmt"
   "strings"
)

func main() {
   s := `PERFORMER "Chris Isaak"
TITLE "Heart Shaped World"
FILE "Chris Isaak - Heart Shaped World.flac" WAVE`

   o := csv.NewReader(strings.NewReader(s))
   o.Comma = ' '
   o.FieldsPerRecord = -1

   for {
      a, e := o.Read()
      if e != nil {
         break
      }
      fmt.Printf("%q\n", a)
   }
}
