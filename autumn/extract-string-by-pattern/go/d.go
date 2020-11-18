package main
import "regexp"

func main() {
   a := regexp.MustCompile("a(..)").FindStringSubmatch("January")
   s := a[1]
   println(s == "nu")
}
