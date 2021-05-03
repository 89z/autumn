package main
import "net/url"

func main() {
   m := url.Values{
      "west": {"left"}, "east": {"right"},
   }
   s := m.Encode()
   println(s == "east=right&west=left")
}
