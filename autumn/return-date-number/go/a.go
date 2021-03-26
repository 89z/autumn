package main
import "time"

func main() {
   t := time.Now()
   n := t.Unix()
   println(n)
}
