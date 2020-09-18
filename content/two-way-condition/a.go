package main

func main() {
   n := 10
   // example A
   sA := ""
   if n > 0 {
      sA = "+"
   } else if n < 0 {
      sA = "-"
   } else {
      sA = "zero"
   }
   // example B
   sB := map[bool]string{true: "+", false: "-"}[n > 0]
   // print
   println(sA == "+", sB == "+")
}
