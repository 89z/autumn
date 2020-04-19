package main
import "fmt"
import "strings"
func main() {
   var o1 strings.Builder
   fmt.Fprint(&o1, "Sunday")
   var s1 = o1.String()
   println(s1)
}
