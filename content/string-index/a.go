package main
import "strings"
func main() {
   s := "Sunday"
   n := strings.Index(s, "day")
   println(n == 3)
}
