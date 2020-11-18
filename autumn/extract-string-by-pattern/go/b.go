package main
import "regexp"

func main() {
   a := regexp.MustCompile("a(..)").FindAllStringSubmatch("January", -1)
   s, s2 := a[0][1], a[1][1]
   println(s == "nu", s2 == "ry")
}
