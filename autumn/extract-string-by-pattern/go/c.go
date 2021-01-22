package main
import "regexp"

func findSubmatch(re, input string) string {
   a := regexp.MustCompile(re).FindStringSubmatch(input)
   if len(a) < 2 {
      return ""
   }
   return a[1]
}

func main() {
   s := findSubmatch("a(..)", "January")
   println(s == "nu")
}
