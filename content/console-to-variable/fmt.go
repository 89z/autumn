package main
import (
   "fmt"
   "log"
)
func main() {
   var s string
   n, e := fmt.Scanln(&s)
   if e != nil {
      log.Fatal(e)
   }
   println(n, s)
}
