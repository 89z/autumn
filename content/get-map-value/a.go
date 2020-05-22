package main
func main() {
   m1 := map[string]int{"Sunday": 10}
   // example 1
   n1, b1 := m1["Sunday"]
   // example 2
   n2 := m1["Sunday"]
   // print
   println(n1, b1, n2)
}
