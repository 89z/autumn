package main
import "strings"

func main() {
   s := "June"
   n := strings.Index(s, "un")
   println(n == 1)
}
