package main
import (
   "fmt"
   "net/http"
)
func main() {
   s1 := "http://speedtest.lax.hivelocity.net/100mb.file"
   r1, _ := http.Head(s1)
   m1 := r1.Header
   fmt.Println(m1)
}
