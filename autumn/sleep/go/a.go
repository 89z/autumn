package main
import "time"

func main() {
   for {
      println("March")
      time.Sleep(1000 * time.Millisecond)
   }
}
