package main
import (
   "io"
   "net/http"
   "os"
)
func main() {
   o_in, e := http.Get("https://speedtest.lax.hivelocity.net")
   if e != nil {
      os.Exit(1)
   }
   o_out, e := os.Create("index.html")
   if e != nil {
      os.Exit(1)
   }
   io.Copy(o_out, o_in.Body)
}
