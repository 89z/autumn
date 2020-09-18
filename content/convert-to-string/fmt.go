package main
import "fmt"

func main() {
   n := 9
   // example 1
   s := fmt.Sprint(n)
   // example 2
   s2 := fmt.Sprintf("%d", n)
   // print
   println(s == "9", s2 == "9")
}
