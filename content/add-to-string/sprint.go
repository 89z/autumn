package main
import "fmt"

func main() {
   s := "May"
   // example 1
   s = fmt.Sprint(s, ",June")
   // example 2
   s = fmt.Sprintf("%s,July", s)
   // print
   println(s == "May,June,July")
}
