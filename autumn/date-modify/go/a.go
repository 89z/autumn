package main
import "time"

func main() {
   fmt.Printf("Yesterday: %v\n", time.Now().Add(-24*time.Hour))
}
