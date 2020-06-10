package main
import (
   "fmt"
   "log"
   "time"
)
func main() {
   s := time.RFC3339[:10]
   o, e := time.Parse(s, "2019-12-31")
   if e != nil {
      log.Fatal(e)
   }
   fmt.Println(o)
}
