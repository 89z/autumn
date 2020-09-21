package main

func main() {
   n := 10
   // example 1
   s1 := ""
   if n > 0 {
      s1 = "+"
   } else if n < 0 {
      s1 = "-"
   } else {
      s1 = "zero"
   }
   // example 2
   s2 := map[bool]string{true: "+", false: "-"}[n > 0]
   // print
   println(s1 == "+", s2 == "+")
}
