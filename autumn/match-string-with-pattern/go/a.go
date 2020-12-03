package main
import "regexp"

func main() {
   o := regexp.MustCompile("..n")
   in_s := "Sunday Monday"
   if o.MatchString(in_s) {
      out_s := o.FindString(in_s)
      println(out_s)
   }
}
