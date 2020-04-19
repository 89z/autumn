package main
import "fmt"
import "net/http"
func main() {
   s1 := "http://speedtest.lax.hivelocity.net"
   // example 1
   o1, _ := http.Get(s1)
   fmt.Println(o1)
   // example 2
   o2 := o1.Body
   fmt.Printf("%#v\n", o2)
}
