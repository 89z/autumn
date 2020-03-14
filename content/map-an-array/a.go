package main
import "fmt"
func main() {
   a1 := []int{10, 11}
   for n1 := range a1 {
      a1[n1] += 12
   }
   fmt.Println(a1)
}
