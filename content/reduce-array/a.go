package main

func main() {
   a := []string{"May", "June"}
   // example 1
   f := func (s_acc, s_cur string) string {
      return s_acc + s_cur
   }
   s := ""
   for n := range a {
      s_cur := a[n]
      s = f(s, s_cur)
   }
   // example 2
   s2 := ""
   for n := range a {
      s2 += a[n]
   }
   // print
   println(s == "MayJune", s2 == "MayJune")
}
