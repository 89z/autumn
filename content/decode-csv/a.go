package main
import (
   "encoding/csv"
   "fmt"
   "log"
   "os"
)
func main() {
   o_db, e := os.Open("a.csv")
   if e != nil {
      log.Fatal(e)
   }
   o_tab := csv.NewReader(o_db)
   a_head, e := o_tab.Read()
   if e != nil {
      log.Fatal(e)
   }
   m_row := map[string]string{}
   for {
      a_row, e := o_tab.Read()
      if e != nil {
         break
      }
      for n_row, s_row := range a_head {
         m_row[s_row] = a_row[n_row]
      }
      fmt.Println(m_row)
   }
}
