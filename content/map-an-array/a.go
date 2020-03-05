package main
import "fmt"
func main() {
   a1 := []int{10, 20}
   for n1 := range a1 {
      a1[n1] += 30
   }
   fmt.Println(a1)
}
