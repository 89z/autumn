package main

import (
   "encoding/csv"
   "fmt"
   "strings"
)

func main() {
   s := `PERFORMER "Ennio Morricone"
TITLE "Very Best of Ennio Morricone"
FILE "Ennio Morricone - Very Best of Ennio Morricone.flac" WAVE
  TRACK 01 AUDIO`

   o := csv.NewReader(strings.NewReader(s))
   o.Comma = ' '
   o.FieldsPerRecord = -1
   o.TrimLeadingSpace = true

   for {
      a, e := o.Read()
      if e != nil {
         break
      }
      fmt.Printf("%q\n", a)
   }
}
