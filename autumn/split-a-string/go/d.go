package main
import "strings"

func main() {
   s := "May June July"
   a := strings.SplitN(s, " ", 2)
   println(a[1] == "June July")
}
