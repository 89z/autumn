package main
import "fmt"

func main() {
   var s = "11"
   var n int
   fmt.Sscanf(s, "%v", &n)
   fmt.Println(n == 11)
}
