package main
import "regexp"

func main() {
   b := []byte("south north")
   c := regexp.MustCompile("o(..)").FindSubmatch(b)[1]
   println(string(c) == "ut")
}
