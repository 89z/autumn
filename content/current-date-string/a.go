package main
import "time"
func main() {
   var o = time.Now()
   var s string
   // example 1
   s = o.String()
   println(s)
   // example 2
   s = o.Format(time.ANSIC)
   println(s)
   // example 3
   s = o.Format("Mon Jan 2 15:04:05 2006")
   println(s)
}
