package main
import (
   "fmt"
   "io/ioutil"
   "net/http"
)
func main() {
   // read
   s1 := "http://speedtest.lax.hivelocity.net"
   o1, _ := http.Get(s1)
   a1, _ := ioutil.ReadAll(o1.Body)
   // write
   fmt.Printf("%s", a1)
}