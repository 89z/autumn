package main
import "strings"
func main() {
   var o1 strings.Builder
   o1.WriteString("Sunday")
   var s1 = o1.String()
   println(s1)
}
