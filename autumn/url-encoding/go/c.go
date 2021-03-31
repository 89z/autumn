package main
import "net/url"

func main() {
   u := url.URL{Host: "golang.org"}
   s := u.String()
   println(s)
}
