package main
import (
   "io/ioutil"
   "log"
   "net/http"
)
func main() {
   s := "https://speedtest.lax.hivelocity.net"
   o, e := http.Get(s)
   if e != nil {
      log.Fatal(e)
   }
   y, e := ioutil.ReadAll(o.Body)
   if e != nil {
      log.Fatal(e)
   }
   log.Printf("%s", y)
}
