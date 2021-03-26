package main
import "strings"

func main() {
   var b strings.Builder
   b.WriteString("March")
   s := b.String()
   println(s == "March")
}
