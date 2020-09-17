package main

func main() {
   a := []string{"May", "June"}
   s := ""
   for n := range a {
      s += a[n]
   }
   println(s == "MayJune")
}
