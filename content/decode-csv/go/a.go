package main

import (
   "encoding/csv"
   "log"
   "strings"
)

func f(s_in string) []map[string]string {
   a_in, e := csv.NewReader(strings.NewReader(s_in)).ReadAll()
   if e != nil {
      log.Fatal(e)
   }
   a_out := []map[string]string{}
   for n_row, a_row := range a_in {
      if n_row == 0 {
         continue
      }
      m_row := map[string]string{}
      for n_col, s_col := range a_in[0] {
         m_row[s_col] = a_row[n_col]
      }
      a_out = append(a_out, m_row)
   }
   return a_out
}

func main() {
   s := `Month,Day
January,Sunday
February,Monday`
   a := f(s)
   log.Print(a)
}
