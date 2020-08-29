package main
import (
   "io"
   "log"
   "net/http"
   "os"
)
func main() {
   o_in, e := http.Get("https://speedtest.lax.hivelocity.net")
   if e != nil {
      log.Fatal(e)
   }
   o_out, e := os.Create("index.html")
   if e != nil {
      log.Fatal(e)
   }
   io.Copy(o_out, o_in.Body)
}
