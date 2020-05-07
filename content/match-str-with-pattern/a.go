package main
import "regexp"
func main() {
   s1 := "Sunday"
   // example 1
   b1 := regexp.MustCompile("^Su").MatchString(s1)
   // example 2
   b2 := regexp.MustCompile("un.").MatchString(s1)
   // example 3
   b3 := regexp.MustCompile("ay$").MatchString(s1)
   // example 4
   b4 := regexp.MustCompile("(?i)su").MatchString(s1)
   // print
   println(b1, b2, b3, b4)
}
