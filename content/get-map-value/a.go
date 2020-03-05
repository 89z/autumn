package main
func main() {
   m1 := map[string]int{"sun": 10, "mon": 11}
   // example 1
   n1, _ := m1["mon"]
   // example 2
   n2 := m1["mon"]
   // print
   println(n1, n2)
}
