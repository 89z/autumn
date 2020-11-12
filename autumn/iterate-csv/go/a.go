package main

import (
   "encoding/csv"
   "log"
   "strings"
)

func main() {
   str_o := strings.NewReader(`Month,Day
January,Sunday
February,Monday`)
   csv_o := csv.NewReader(str_o)
   head_a, e := csv_o.Read()
   if e != nil {
      log.Fatal(e)
   }
   for {
      row_a, e := csv_o.Read()
      if e != nil {
         break
      }
      m := map[string]string{}
      for n, s := range head_a {
         m[s] = row_a[n]
      }
      log.Print(m)
   }
}
