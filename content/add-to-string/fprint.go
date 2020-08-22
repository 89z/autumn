package main
import (
   "fmt"
   "strings"
)
func main() {
   o := strings.Builder{}
   fmt.Fprint(&o, "Sunday")
   s := o.String()
   println(s)
}
