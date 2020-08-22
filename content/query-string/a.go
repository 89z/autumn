package main
import (
   "fmt"
   "net/url"
   "os"
)
func main() {
   // example 1
   m1, e := url.ParseQuery("one=odd&two=even")
   if e != nil {
      os.Exit(1)
   }
   // example 2
   m2 := url.Values{
      "one": []string{"odd"},
      "two": []string{"even"},
   }
   s1 := m2.Encode()
   // print
   fmt.Println(m1, s1)
}
