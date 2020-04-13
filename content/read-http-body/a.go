package main
import (
   "fmt"
   "io/ioutil"
   "net/http"
)
func main() {
   s1 := "http://speedtest.lax.hivelocity.net"
   r1, _ := http.Get(s1)
   a1, _ := ioutil.ReadAll(r1.Body)
   fmt.Printf("%s", a1)
}
