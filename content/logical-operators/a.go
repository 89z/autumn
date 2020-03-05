package main
func main() {
   n1 := 10
   // not
   b1 := !(n1 > 20)
   // or
   b2 := n1 > 30 || n1 < 20
   // and
   b3 := n1 > 0 && n1 < 40
   // print
   println(b1, b2, b3)
}
