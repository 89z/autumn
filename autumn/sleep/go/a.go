package main
import "time"

func main() {
   for {
      println("Sleep")
      time.Sleep(500 * time.Millisecond)
   }
}
