package main
import "net/url"
func main() {
   s1 := "https://example.com/one?two=even"
   o1, _ := url.Parse(s1)
   // example 1
   s2 := o1.Host
   // example 2
   s3 := o1.Path
   // example 3
   s4 := o1.RawQuery
   // print
   println(s2 == "example.com", s3 == "/one", s4 == "two=even")
}
