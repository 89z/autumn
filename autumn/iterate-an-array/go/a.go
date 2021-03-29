package main

func main() {
   a := []string{"west", "east"}
   for n := 0; n < len(a); n++ {
      s := a[n]
      println(n, s)
   }
}
