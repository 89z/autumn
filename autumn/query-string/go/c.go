package main
import "net/url"

func main() {
   m := url.Values{}
   m.Set("one", "odd")
   m.Set("two", "even")
   s := m.Encode()
   println(s == "one=odd&two=even")
}
