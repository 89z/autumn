package main
import "regexp"

func main() {
   s := "January"
   b := regexp.MustCompile("(?i)ja").MatchString(s)
   println(b)
}
