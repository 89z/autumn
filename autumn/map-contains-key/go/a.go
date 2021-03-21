package main

func main() {
   m := map[string]int{"month": 5, "day": 4}
   { // example 1
      n, b := m["day"]
      if b {
         println(n)
      }
   }
   { // example 2
      n := m["day"]
      if n != 0 {
         println(n)
      }
   }
}
