package main
func main() {
   m1 := map[string]int{"Sun": 10, "Mon": 11}
   // example 1
   n1, _ := m1["Mon"]
   // example 2
   n2 := m1["Mon"]
   // print
   println(n1, n2)
}
