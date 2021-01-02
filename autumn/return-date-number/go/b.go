package main
import "time"

func main() {
   o := time.Now()
   n := o.UnixNano() / 1e9
   println(n)
}
