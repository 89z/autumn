package main
import "regexp"

func main() {
   b := []byte("south north")
   c := regexp.MustCompile("o..").Find(b)
   println(string(c) == "out")
}
