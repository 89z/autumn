package main
import (
   "fmt"
   "log"
   "net/url"
)
func main() {
   // example 1
   s1 := "one=odd&two=even"
   m1, e1 := url.ParseQuery(s1)
   if e1 != nil {
      log.Fatal(e1)
   }
   // example 2
   m2 := url.Values{
      "one": []string{"odd"},
      "two": []string{"even"},
   }
   s2 := m2.Encode()
   // print
   fmt.Println(m1, s2)
}
