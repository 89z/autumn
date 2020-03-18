package main
func main() {
   a1 := []string{"Sun", "Mon"}
   // example 1
   for _, s1 := range a1 {
      println(s1)
   }
   // example 2
   for n1, s1 := range a1 {
      println(n1, s1)
   }
   // example 3
   var n1 int
   var s1 string
   for n1, s1 = range a1 {
      println(n1, s1)
   }
}
