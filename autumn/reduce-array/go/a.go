package main

func fold(a []string, f func(string, string) string) string {
   var s string
   for _, t := range a {
      s = f(s, t)
   }
   return s
}

func main() {
   a := []string{"May", "June"}
   s := fold(a, func (s, t string) string {
      return s + t
   })
   println(s == "MayJune")
}
