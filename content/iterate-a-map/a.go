package main
func main() {
   m1 := map[string]int{"sun": 10, "mon": 11}
   // example 1
   for s1 := range m1 {
      println(s1)
   }
   // example 2
   for s1, n1 := range m1 {
      println(s1, n1)
   }
}
