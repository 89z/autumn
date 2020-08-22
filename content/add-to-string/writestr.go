package main
import "strings"
func main() {
   o := strings.Builder{}
   o.WriteString("Sunday")
   s := o.String()
   println(s)
}
