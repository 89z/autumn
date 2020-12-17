package main
import "regexp"

func main() {
   in_s := "January"
   o := regexp.MustCompile("a.")
   if o.MatchString(in_s) {
      out_s := o.FindString(in_s)
      println(out_s)
   }
}
