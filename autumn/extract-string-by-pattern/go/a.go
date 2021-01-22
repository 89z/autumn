package main
import "regexp"

func find(re string, input []byte) string {
   b := regexp.MustCompile(re).Find(input)
   return string(b)
}

func main() {
   s := find("a..", []byte("January"))
   println(s == "anu")
}
