package main
import "time"

func main() {
   println("May")
   o := time.Duration(5 * time.Second)
   time.Sleep(o)
}
