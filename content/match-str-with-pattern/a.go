package main
import "regexp"
func main() {
   s := "Wednesday"
   // example 1
   b1 := regexp.MustCompile("^W").MatchString(s)
   // example 2
   b2 := regexp.MustCompile("(?i)we").MatchString(s)
   // print
   println(b1, b2)
}
