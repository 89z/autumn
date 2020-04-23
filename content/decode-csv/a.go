package main
import (
   "encoding/csv"
   "os"
)
func main() {
   o1, _ := os.Open("a.csv")
   o2 := csv.NewReader(o1)
   m1 := map[string]int{}
   for {
      a1, e1 := o2.Read()
      if e1 != nil {
         break
      }
      if len(m1) == 0 {
         for n1, s1 := range a1 {
            m1[s1] = n1
         }
         continue
      }
      s1 := a1[m1["city"]]
      println(s1)
   }
}
