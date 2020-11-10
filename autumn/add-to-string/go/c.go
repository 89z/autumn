package main
import "strings"

func main() {
   o := strings.Builder{}
   o.WriteString("March")
   s := o.String()
   println(s == "March")
}
