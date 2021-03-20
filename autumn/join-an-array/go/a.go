package main
import "strings"

func main() {
   // example 1
   a1 := []string{"May", "June"}
   s1 := strings.Join(a1, ",")
   // example 2
   a2 := []string{"March"}
   s2 := strings.Join(a2, ",")
   // print
   println(s1 == "May,June", s2 == "March")
}
