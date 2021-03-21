package main
import "strings"

func main() {
   { // example 1
      a := []string{"May", "June"}
      s := strings.Join(a, ",")
      println(s == "May,June")
   }
   { // example 2
      a := []string{"March"}
      s := strings.Join(a, ",")
      println(s == "March")
   }
}
