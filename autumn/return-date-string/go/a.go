package main
import "time"

func main() {
   s := time.Now().Format(time.RFC3339)
   println(s)
}
