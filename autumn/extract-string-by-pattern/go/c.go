package main
import "regexp"

func main() {
   s := "south north"
   t := regexp.MustCompile("o(..)").FindStringSubmatch(s)[1]
   println(t == "ut")
}
