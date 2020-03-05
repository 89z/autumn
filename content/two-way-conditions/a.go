package main
func main() {
   var n1 = 10
   var n2 int
   if n1 > 30 {
      n2 = 30
   } else if n1 > 20 {
      n2 = 20
   } else {
      n2 = n1
   }
   println(n2)
}
