package main
import "fmt"

func main() {
   n := 11
   // example 1
   s1 := fmt.Sprintf("%d", n)
   // example 2
   s2 := fmt.Sprintf("%v", n)
   // print
   fmt.Println(s1 == "11", s2 == "11")
}
