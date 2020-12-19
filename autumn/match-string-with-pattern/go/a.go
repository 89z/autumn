package main
import "regexp"

func main() {
   b := regexp.MustCompile("a..").MatchString("January")
   println(b)
}
