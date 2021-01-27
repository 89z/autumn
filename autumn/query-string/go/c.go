package main
import "net/url"

func main() {
   v := url.Values{
      "one": []string{"odd"},
      "two": []string{"even"},
   }
   // example 1
   s1 := v.Get("one")
   // example 2
   s2 := v["one"][0]
   // print
   println(s1 == "odd", s2 == "odd")
}
