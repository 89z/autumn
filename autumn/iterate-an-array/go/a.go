package main

func main() {
   s := []string{"west", "east"}
   for n := 0; n < len(s); n++ {
      st := s[n]
      println(n, st)
   }
}
