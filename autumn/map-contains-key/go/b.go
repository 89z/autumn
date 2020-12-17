package main

func main() {
   m := map[string]interface{}{"month": 12, "day": 31}
   // example 1
   n1, b := m["day"]
   if b {
      println(n1.(int))
   }
   // example 2
   n2 := m["day"]
   if n2 != nil {
      println(n2.(int))
   }
}
