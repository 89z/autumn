package main

func main() {
   m := map[string]int{"month": 12, "day": 31}
   n, b := m["day"]
   println(n == 31, b)
}
