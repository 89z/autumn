package main
import (
   "fmt"
   "net/url"
   "os"
)
func main() {
   o1, e := url.Parse("https://example.com/one?two=even")
   if e != nil {
      os.Exit(1)
   }
   o2 := o1.Query()
   fmt.Println(o2)
}
