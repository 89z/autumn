package main

func main() {
   m := map[string]int{"month": 12, "day": 31}
   // example 1
   n1, b1 := m["day"]
   // example 2
   n2, b2 := m["year"]
   // print
   println(n1 == 31, b1, n2 == 0, ! b2)
}
