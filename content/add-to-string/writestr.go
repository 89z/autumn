package main
import "strings"

func main() {
   o := strings.Builder{}
   o.WriteString("May")
   s := o.String()
   println(s == "May")
}
