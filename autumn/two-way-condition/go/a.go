package main

func main() {
   var (
      n = 10
      s string
   )

   // example 1
   if n > 0 {
      s = "+"
   } else if n < 0 {
      s = "-"
   } else {
      s = "zero"
   }
   println(s == "+")

   // example 2
   s = map[bool]string{true: "+", false: "-"}[n > 0]
   println(s == "+")
}
