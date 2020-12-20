package main
import "fmt"

func IntVal(s string) int {
   var n int
   fmt.Sscan(s, &n)
   return n
}

func main() {
   n := IntVal("100")
   fmt.Println(n)
}
