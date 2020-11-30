package main

func main() {
   m := map[string]int{"month": 12, "day": 31}

   println("example 1")
   for s := range m {
      println(s)
   }

   println("example 2")
   for s, n := range m {
      println(s, n)
   }
}
