package main
import "regexp"

func main() {
   s := "January"
   // example 1
   b := regexp.MustCompile("^J").MatchString(s)
   // example 2
   b2 := regexp.MustCompile("(?i)ja").MatchString(s)
   // print
   println(b, b2)
}
