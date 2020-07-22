package main
import (
   "io"
   "log"
   "net/http"
   "os"
)
func main() {
   o1, e := http.Get("http://speedtest.lax.hivelocity.net")
   if e != nil {
      log.Fatal(e)
   }
   o2, e := os.Create("index.html")
   if e != nil {
      log.Fatal(e)
   }
   io.Copy(o2, o1.Body)
}
