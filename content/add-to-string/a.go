package main
import "fmt"
func main() {
   // example 1
   var s1 = "Sun"
   s1 = s1 + "day"
   // example 2
   var s2 = "Sun"
   s2 += "day"
   // example 3
   var s3 = "Sun"
   s3 = fmt.Sprintf("%s Mon", s3)
   // print
   fmt.Printf("%q\n", []string{s1, s2, s3})
}
