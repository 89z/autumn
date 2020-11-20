package main
import "fmt"

func main() {
   var n int
   fmt.Sscanln("100", &n)
   fmt.Println(n == 100)
}
