package main

func main() {
   s := []string{"west", "east"}
   // example 1
   for n, st := range s {
      println(n, st)
   }
   // example 2
   for n := range s {
      st := s[n]
      println(n, st)
   }
}
