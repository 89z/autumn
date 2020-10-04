package main
import "strconv"

func main() {
   n := 9.9
   // example 1
   s1 := strconv.FormatFloat(n, 'f', -1, 64)
   // example 2
   s2 := strconv.FormatFloat(n, 'f', 0, 64)
   // print
   println(s1 == "9.9", s2 == "10")
}
