package main

func main() {
   m := map[string]interface{}{"month": 5, "day": 4}
   { // example 1
      n, b := m["day"]
      if b {
         println(n.(int))
      }
   }
   { // example 2
      n := m["day"]
      if n != nil {
         println(n.(int))
      }
   }
}
