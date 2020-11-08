package main

func FoldIntString(a []string, f func(int, string) int) int {
   n := 0
   for _, s := range a {
      n = f(n, s)
   }
   return n
}

func main() {
   a := []string{"May", "June"}
   n := FoldIntString(a, func (n int, s string) int {
      return n + len(s)
   })
   println(n == 7)
}
