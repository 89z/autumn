package main
import "fmt"

func main() {
   var s = "11"
   var n int
   fmt.Sscanln(s, &n)
   fmt.Println(n == 11)
}
