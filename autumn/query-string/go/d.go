package main
import "net/url"

func main() {
   v := make(url.Values)
   v.Set("one", "odd")
   v.Set("two", "even")
   s := v.Encode()
   println(s == "one=odd&two=even")
}
