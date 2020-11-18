package main
import "regexp"

func main() {
   a := regexp.MustCompile("a.").FindAllString("January", -1)
   s, s2 := a[0], a[1]
   println(s == "an", s2 == "ar")
}
