package main
import "fmt"
func main() {
   a := []string{"Sun", "Mon", "Tue"}
   // string from front
   s := a[0]
   // string from back
   s2 := a[len(a) - 1]
   // array from front
   a2 := a[:2]
   // array from back
   a3 := a[len(a) - 2:]
   // print
   fmt.Println(s, s2, a2, a3)
}
