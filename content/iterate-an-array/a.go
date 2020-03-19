package main
func main() {
   a1 := []string{"Sun", "Mon"}
   // example 1
   for n1, s1 := range a1 {
      println(n1, s1)
   }
   // example 2
   for n1 := range a1 {
      println(n1)
   }
   // example 3
   for _, s1 := range a1 {
      println(s1)
   }
}
