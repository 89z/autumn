package main
import "net/url"

func main() {
   m := url.Values{
      "one": []string{"odd"},
      "two": []string{"even"},
   }
   s := m.Encode()
   println(s == "one=odd&two=even")
}
