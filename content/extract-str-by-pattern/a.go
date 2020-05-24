package main
import (
   "fmt"
   "regexp"
)
func main() {
   s1 := "Wednesday"
   s2 := "e."
   // example 1
   s3 := regexp.MustCompile(s2).FindString(s1)
   // example 2
   a1 := regexp.MustCompile(s2).FindAllString(s1, -1)
   // print
   fmt.Println(s3, a1)
}
