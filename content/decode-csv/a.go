package main
import (
   "encoding/csv"
   "os"
)
func main() {
   r1, _ := os.Open("a.csv")
   r2 := csv.NewReader(r1)
   m1 := map[string]int{}
   for {
      a1, e1 := r2.Read()
      if e1 != nil {
         break
      }
      if len(m1) == 0 {
         for n1, s1 := range a1 {
            m1[s1] = n1
         }
         continue
      }
      println(a1[m1["city"]])
   }
}
