package main
import "regexp"

func main() {
   s := "January"
   b := regexp.MustCompile("^J").MatchString(s)
   println(b)
}
