package main

func main() {
   a := []string{"May", "June"}
   for n := 0; n < len(a); n++ {
      s := a[n]
      println(n, s)
   }
}
