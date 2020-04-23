package main
import (
   "fmt"
   "net/http"
)
func main() {
   s1 := "http://speedtest.lax.hivelocity.net/100mb.file"
   o1, _ := http.Head(s1)
   fmt.Println(o1)
}
