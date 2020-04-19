package main
import "fmt"
func main() {
   s1 := "Sunday"
   // example 1
   s2 := fmt.Sprint(s1, "Monday")
   // example 2
   s3 := fmt.Sprintf("%sMonday", s1)
   // print
   println(s2, s3)
}
