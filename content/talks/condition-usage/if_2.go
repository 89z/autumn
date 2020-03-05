package main
func main() {
   var n1 = 10
   var n2 int
   if n1 { // non-bool n1 (type int) used as if condition
      n2 = n1
   } else {
      n2 = 20
   }
   println(n2)
}
