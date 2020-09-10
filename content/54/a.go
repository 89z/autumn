package main
import (
   "fmt"
   "io/ioutil"
   "log"
   "net/http"
)
func main() {
   o, e := http.Get("https://speedtest.lax.hivelocity.net")
   if e != nil {
      log.Fatal(e)
   }
   y, e := ioutil.ReadAll(o.Body)
   if e != nil {
      log.Fatal(e)
   }
   fmt.Printf("%s", y)
}
