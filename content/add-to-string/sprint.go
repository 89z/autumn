package main
import "fmt"

func main() {
   s := "May"
   // example 1
   s1 := fmt.Sprint(s, ",June")
   // example 2
   s2 := fmt.Sprintf("%s,June", s)
   // print
   println(s1 == "May,June", s2 == "May,June")
}
