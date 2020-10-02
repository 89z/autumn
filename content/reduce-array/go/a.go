package main

func f(s, s1 string) string {
   return s + s1
}

func main() {
   a := []string{"May", "June"}
   s := ""
   for n := range a {
      s1 := a[n]
      s = f(s, s1)
   }
   println(s == "MayJune")
}
