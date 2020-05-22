package main
import (
   "log"
   "net/url"
)
func main() {
   s1 := "https://example.com/one?two=even"
   o1, e := url.Parse(s1)
   if e != nil {
      log.Fatal(e)
   }
   o2 := o1.Query()
   log.Print(o2)
}
