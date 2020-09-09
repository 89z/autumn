package main
func main() {
   n := 1
   // example 1
   s := ""
   if n > 0 {
      s = "+"
   } else if n < 0{
      s = "-"
   } else {
      s = "zero"
   }
   // example 2
   s2 := map[bool]string{true: "+", false: "-"}[n > 0]
   // print
   println(s == "+", s2 == "+")
}
