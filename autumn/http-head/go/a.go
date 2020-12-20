package main
import "net/http"

func HttpHead(s string) bool {
   o, e := http.Head(s)
   return e == nil && o.StatusCode == 200
}

func main() {
   b := HttpHead("http://speedtest.lax.hivelocity.net")
   println(b)
}
