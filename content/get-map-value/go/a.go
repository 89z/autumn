package main

func main() {
   m := map[string]int{"year": 2019}
   // example 1
   n1, b1 := m["year"]
   // example 2
   n2 := m["year"]
   // print
   println(b1, n1 == 2019, n2 == 2019)
}
