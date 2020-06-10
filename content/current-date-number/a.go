package main
import "time"
func main() {
   o := time.Now()
   n := o.Unix()
   println(n)
}
