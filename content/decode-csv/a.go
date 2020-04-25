package main
import (
   "encoding/csv"
   "os"
)
func main() {
   o_open, _ := os.Open("a.csv")
   o_table := csv.NewReader(o_open)
   a_head, _ := o_table.Read()
   m_head := map[string]int{}
   for n_ind, s_col := range a_head {
      m_head[s_col] = n_ind
   }
   a_body, _ := o_table.ReadAll()
   for _, a_row := range a_body {
      n_city := m_head["city"]
      s_city := a_row[n_city]
      println(s_city)
   }
}
