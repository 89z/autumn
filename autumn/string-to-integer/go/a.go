package main
import "fmt"

func intVal(s string) int {
   var n int
   fmt.Sscan(s, &n)
   return n
}

func main() {
   n := intVal("100")
   fmt.Println(n)
}
