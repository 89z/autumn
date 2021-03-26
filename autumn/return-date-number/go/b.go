package main
import "time"

func main() {
   t := time.Now()
   n := t.UnixNano() / 1e9
   println(n)
}
