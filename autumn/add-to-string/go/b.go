package main
import "fmt"

func main() {
   s := "March"
   // example 1
   s1 := fmt.Sprintf("%v April", s)
   // example 2
   s2 := fmt.Sprintf("%v %[1]v", s)
   // print
   fmt.Println(s1 == "March April", s2 == "March March")
}
