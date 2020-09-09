package main
import "strings"

func main() {
   a := []string{"May", "June"}
   s := strings.Join(a, ",")
   println(s == "May,June")
}
