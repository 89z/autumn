package main
import "reflect"
func main() {
   a1 := []int{10, 11}
   a2 := []int{10, 11}
   b1 := reflect.DeepEqual(a1, a2)
   println(b1)
}
