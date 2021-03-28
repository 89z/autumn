package main
import "strings"

func main() {
   s := "north"
   b := strings.HasPrefix(s, "no")
   println(b)
}
