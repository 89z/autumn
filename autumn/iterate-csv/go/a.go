package main

import (
   "encoding/csv"
   "log"
   "strings"
)

func main() {
   o_str := strings.NewReader(`Month,Day
January,Sunday
February,Monday`)
   o_csv := csv.NewReader(o_str)
   a_head, e := o_csv.Read()
   if e != nil {
      log.Fatal(e)
   }
   for {
      a_row, e := o_csv.Read()
      if e != nil {
         break
      }
      m := map[string]string{}
      for n, s := range a_head {
         m[s] = a_row[n]
      }
      log.Print(m)
   }
}
