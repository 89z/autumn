package main
import (
   "encoding/csv"
   "fmt"
   "log"
   "os"
)
func main() {
   o_csv, e := os.Open("a.csv")
   if e != nil {
      log.Fatal(e)
   }
   o_table := csv.NewReader(o_csv)
   a_head, e := o_table.Read()
   if e != nil {
      log.Fatal(e)
   }
   a_body, e := o_table.ReadAll()
   if e != nil {
      log.Fatal(e)
   }
   m_row := map[string]string{}
   for n_row, a_row := range a_body {
      for n_col, s_col := range a_head {
         m_row[s_col] = a_row[n_col]
      }
      fmt.Println(n_row, m_row)
   }
}
