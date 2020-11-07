package main
import "regexp"

func main() {
   s := regexp.MustCompile("a.").FindString("January")
   println(s)
}
