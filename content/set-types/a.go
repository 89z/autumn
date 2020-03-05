package main
import "fmt"
func main() {
   // duplicate key 20 in map literal
   c1 := map[int]bool{10: true, 20: true}
   fmt.Println(c1)
}
