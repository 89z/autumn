package main
import "strconv"
func main() {
   s1 := "10"
   // example 1
   n1, _ := strconv.ParseInt(s1, 10, 0)
   // example 2
   n2, _ := strconv.Atoi(s1)
   // print
   println(n1, n2)
}
