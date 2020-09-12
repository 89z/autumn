package main
import "strings"

func main() {
   s := "June"
   // example 1
   b := strings.HasPrefix(s, "Ju")
   // example 2
   b2 := strings.Contains(s, "un")
   // example 3
   b3 := strings.HasSuffix(s, "ne")
   // print
   println(b, b2, b3)
}
