package main

func main() {
   a := []string{"May", "June"}
   // example 1
   for n, s := range a {
      println(n, s)
   }
   // example 2
   for n := range a {
      s := a[n]
      println(n, s)
   }
}
