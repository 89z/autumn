package main
import (
   "log"
   "net/url"
)
func main() {
   s1 := "https://example.com/one?two=even"
   o1, e1 := url.Parse(s1)
   if e1 != nil {
      log.Fatal(e1)
   }
   o2 := o1.Query()
   log.Print(o2)
}
