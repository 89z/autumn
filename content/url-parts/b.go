package main
import (
   "fmt"
   "net/url"
)
func main() {
   s1 := "https://example.com/one?two=even"
   o1, _ := url.Parse(s1)
   o2 := o1.Query()
   fmt.Println(o2)
}
