package main
import "net/url"

func main() {
   u := url.URL{Scheme:"https", Host:"musicbrainz.org", Path:"/ws/2/release"}
   s := u.String()
   println(s)
}
