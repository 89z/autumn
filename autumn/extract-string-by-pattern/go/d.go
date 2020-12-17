package main
import "regexp"

func Find(pattern, subject string) string {
   a := regexp.MustCompile(pattern).FindStringSubmatch(subject)
   if len(a) < 2 {
      return ""
   }
   return a[1]
}

func main() {
   s := Find("(..)n", "January")
   println(s == "Ja")
}
