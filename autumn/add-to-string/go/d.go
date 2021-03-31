package main
import "strings"

func main() {
   var b strings.Builder
   b.WriteString("north")
   s := b.String()
   println(s == "north")
}
