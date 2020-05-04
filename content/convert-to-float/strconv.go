package main
import "strconv"
func main() {
   s1 := "1.9"
   n1, _ := strconv.ParseFloat(s1, 64)
   println(n1)
}
