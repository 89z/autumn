package main
import "regexp"

func main() {
   s := "south north"
   b := regexp.MustCompile("o..").MatchString(s)
   println(b)
}
