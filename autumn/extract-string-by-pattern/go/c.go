package main
import "regexp"

func FindSubmatch(pat, sub string) string {
   a := regexp.MustCompile(pat).FindStringSubmatch(sub)
   if len(a) < 2 {
      return ""
   }
   return a[1]
}

func main() {
   s := FindSubmatch("a(..)", "January")
   println(s == "nu")
}
