package main
import "strconv"

func main() {
   // example 1
   n := int64(9)
   s := strconv.FormatInt(n, 10)
   // example 2
   n2 := 9.9
   s2 := strconv.FormatFloat(n2, 'f', -1, 64)
   // print
   println(s == "9", s2 == "9.9")
}
