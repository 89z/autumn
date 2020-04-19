package main
import "fmt"
func main() {
   // example 1
   s1 := "Sun"
   s1 = fmt.Sprint(s1, "Mon")
   // example 2
   s2 := "Sun"
   s2 = fmt.Sprintf("%sMon", s2)
   // print
   println(s1, s2)
}
