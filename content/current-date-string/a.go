package main
import "time"
func main() {
   var o = time.Now()
   var s string
   // example 1
   s = o.String()
   println(s)
   // example 2
   s = o.Format(time.RFC3339)
   println(s)
}
