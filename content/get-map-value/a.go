package main
func main() {
   m := map[string]int{"Sunday": 10}
   // example 1
   n1, b := m["Sunday"]
   // example 2
   n2 := m["Sunday"]
   // print
   println(n1, b, n2)
}
