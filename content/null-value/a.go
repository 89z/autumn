package main
func main() {
   // example 1
   var n1 []int
   // example 2
   var n2 map[string]int
   // example 3
   var n3 interface{}
   // example 4
   var n4 func()
   // print
   println(n1 == nil, n2 == nil, n3 == nil, n4 == nil)
}
