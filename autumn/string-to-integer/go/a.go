package main
import "fmt"

func main() {
   var s = "11"
   var n int
   fmt.Sscan(s, &n)
   fmt.Println(n == 11)
}
