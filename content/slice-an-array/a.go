package main
import "fmt"
func main() {
   a := []string{"Sun", "Mon", "Tue"}
   // string from front
   s1 := a[0]
   // string from back
   s2 := a[len(a) - 1]
   // array from front
   a1 := a[:2]
   // array from back
   a2 := a[len(a) - 2:]
   // print
   fmt.Println(s1, s2, a1, a2)
}
