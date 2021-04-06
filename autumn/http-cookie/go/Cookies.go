package main
import "net/http"

func main() {
   r, e := http.Get("https://stackoverflow.com")
   if e != nil {
      panic(e)
   }
   for _, c := range r.Cookies() {
      println(c.Name, c.Value)
   }
}
