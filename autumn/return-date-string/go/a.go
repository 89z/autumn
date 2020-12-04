package main
import "time"

func main() {
   o := time.Now()
   s := o.Format(time.RFC3339)
   println(s)
}
