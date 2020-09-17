package main

func f(sa, sc string) string {
   return sa + sc
}

func main() {
   a := []string{"May", "June"}
   s := ""
   for n := range a {
      sc := a[n]
      s = f(s, sc)
   }
   println(s == "MayJune")
}
