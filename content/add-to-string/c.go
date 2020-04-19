package main
import "fmt"
func main() {
   // example 1
   s1 := "Sun"
   s1 += "Mon"
   // example 2
   s2 := "Sun"
   s2 = fmt.Sprint(s2, "Mon")
   // example 3
   s3 := "Sun"
   s3 = fmt.Sprintf("%sMon", s3)
   // print
   fmt.Printf("%q\n", []string{s1, s2, s3})
}
