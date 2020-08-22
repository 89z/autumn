package main
import (
   "net/url"
   "os"
)
func main() {
   o, e := url.Parse("https://example.com/one?two=even")
   if e != nil {
      os.Exit(1)
   }
   // example 1
   s1 := o.Host
   // example 2
   s2 := o.Path
   // example 3
   s3 := o.RawQuery
   // print
   println(s1 == "example.com", s2 == "/one", s3 == "two=even")
}
