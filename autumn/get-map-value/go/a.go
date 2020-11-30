package main

func main() {
   m := map[string]int{"month": 12, "day": 31}
   // example 1
   n1, b1 := m["day"]
   // example 2
   n2 := m["day"]
   // print
   println(b1, n1 == 31, n2 == 31)
}
