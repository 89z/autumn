package main
import "regexp"

func findSubmatch(re string, input []byte) string {
   a := regexp.MustCompile(re).FindSubmatch(input)
   if len(a) < 2 {
      return ""
   }
   return string(a[1])
}

func main() {
   s := findSubmatch("a(..)", []byte("January"))
   println(s == "nu")
}
