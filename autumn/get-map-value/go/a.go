package main

func main() {
   m := map[string]int{"month": 12, "day": 31}
   { // example 1
      n, b := m["day"]
      println(n == 31, b)
   }
   { // example 2
      n := m["day"]
      println(n == 31)
   }
}
