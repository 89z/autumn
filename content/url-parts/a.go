package main
import (
   "fmt"
   "net/url"
)
func main() {
   s1 := "http://sun.com/mon?tue=10"
   m1, _ := url.Parse(s1)
   fmt.Printf("%#v\n", m1)
}
