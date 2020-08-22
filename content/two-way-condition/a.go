package main
func main() {
   n := 10
   // example 1
   s1 := ""
   if n > 12 {
      s1 = "Tue"
   } else if n > 11 {
      s1 = "Mon"
   } else {
      s1 = "Sun"
   }
   // example 2
   s2 := map[bool]string{true: "Mon", false: "Sun"}[n > 11]
   // print
   println(s1, s2)
}
