package main

func main() {
   m := map[string]int{"month": 5, "day": 4}
   { // example 1
      n, b := m["day"]
      println(n == 4, b)
   }
   { // example 2
      n := m["day"]
      println(n == 4)
   }
}
