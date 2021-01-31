package main
import "net/url"

func main() {
   v := url.Values{
      "one": []string{"odd"},
      "two": []string{"even"},
   }
   s := v.Encode()
   println(s == "one=odd&two=even")
}
