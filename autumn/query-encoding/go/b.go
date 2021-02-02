package main
import "net/url"

func main() {
   m := url.Values{
      "month": []string{"March"},
      "day": []string{"Friday"},
   }
   s := m.Encode()
   println(s == "day=Friday&month=March")
}
