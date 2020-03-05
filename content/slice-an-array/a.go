package main
import "fmt"
func main() {
   a1 := []string{"sun", "mon", "tue"}
   // string from front
   s1 := a1[0]
   // string from back
   s2 := a1[len(a1) - 1]
   // array from front
   a2 := a1[:2]
   // array from back
   a3 := a1[len(a1) - 2:]
   // print
   fmt.Println(s1, s2, a2, a3)
}
