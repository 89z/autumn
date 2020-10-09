package main

func FoldIntString(a []string, f func(int, string) int) int {
   n := 0
   for _, s := range a {
      n = f(n, s)
   }
   return n
}

func FoldString(a []string, f func(string, string) string) string {
   s := ""
   for _, s1 := range a {
      s = f(s, s1)
   }
   return s
}

func main() {
   a := []string{"May", "June"}
   // example 1
   s := FoldString(a, func (s, s1 string) string {
      return s + s1
   })
   // example 2
   n := FoldIntString(a, func (n int, s string) int {
      return n + len(s)
   })
   // print
   println(n == 7, s == "MayJune")
}
