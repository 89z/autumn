package main

func main() {
   a := []string{"May", "June", "July"}
   s := ""
   for n := range a {
      s += a[n]
   }
   println(s)
}
