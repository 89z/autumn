package main
import "net/http"

func head(s string) bool {
   r, e := http.Head(s)
   return e == nil && r.StatusCode == 200
}

func main() {
   b := head("http://speedtest.lax.hivelocity.net")
   println(b)
}
