package main
import "net/http"

func main() {
   r, e := http.Get("https://stackoverflow.com")
   if e != nil {
      panic(e)
   }
   s := r.Header.Get("Set-Cookie")
   println(s)
}
