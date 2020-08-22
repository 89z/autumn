package main
func main() {
   // example 1
   n1 := len("📕")
   // example 2
   a1 := []rune("📕")
   n2 := len(a1)
   // print
   println(n1 == 4, n2 == 1)
}
