package main

func main() {
   m := map[string]int{"year": 2019}
   n, b := m["year"]
   println(n == 2019, b)
}
