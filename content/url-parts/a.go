package main
import (
   "fmt"
   "net/url"
)
func main() {
   s1 := "https://example.com/one?two=even"
   m1, _ := url.Parse(s1)
   fmt.Printf("%#v\n", m1)
}
