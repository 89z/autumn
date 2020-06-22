package main
func main() {
   a := []string{"Sunday", "Monday"}
   // example 1
   for n, s := range a {
      println(n, s)
   }
   // example 2
   for n := range a {
      println(n)
   }
}
