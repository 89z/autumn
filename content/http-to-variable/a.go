package main
import (
   "fmt"
   "io/ioutil"
   "net/http"
   "os"
)
func main() {
   o, e := http.Get("https://speedtest.lax.hivelocity.net")
   if e != nil {
      os.Exit(1)
   }
   y, e := ioutil.ReadAll(o.Body)
   if e != nil {
      os.Exit(1)
   }
   fmt.Printf("%s", y)
}
