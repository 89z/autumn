package main
import "strings"

func main() {
   s := "March"
   b := strings.HasPrefix(s, "Ma")
   println(b)
}
