package main
import (
   "fmt"
   "strings"
)
func main() {
   var o strings.Builder
   fmt.Fprint(&o, "Sunday")
   var s = o.String()
   println(s)
}
