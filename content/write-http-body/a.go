package main
import (
   "io"
   "net/http"
   "os"
)
func main() {
   s1 := "http://speedtest.lax.hivelocity.net"
   o1, _ := http.Get(s1)
   o2, _ := os.Create("a.html")
   io.Copy(o2, o1.Body)
}
