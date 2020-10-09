package main

func FoldString(a []string, f func(string, string) string) string {
   s := ""
   for _, s1 := range a {
      s = f(s, s1)
   }
   return s
}

func main() {
   a := []string{"May", "June"}
   s := FoldString(a, func (s, s1 string) string {
      return s + s1
   })
   println(s == "MayJune")
}
