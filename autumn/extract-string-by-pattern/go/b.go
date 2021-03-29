package main
import "regexp"

func main() {
   s := regexp.MustCompile("o..").FindString("south north")
   println(s == "out")
}
