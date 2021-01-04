package main

func main() {
   m := map[string]int{"month": 5, "day": 4}
   // example 1
   n1, b := m["day"]
   if b {
      println(n1)
   }
   // example 2
   n2 := m["day"]
   if n2 != 0 {
      println(n2)
   }
}
