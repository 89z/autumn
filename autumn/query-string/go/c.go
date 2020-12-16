package main
import "net/url"

func main() {
   m := url.Values{
      "one": []string{"odd"},
      "two": []string{"even"},
   }
   // example 1
   s1 := m.Get("one")
   // example 2
   s2 := m["one"][0]
   // print
   println(s1 == "odd", s2 == "odd")
}
