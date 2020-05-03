package main
import "fmt"
func main() {
   a1 := []string{"Sun", "Mon"}
   a2 := []int{10, 11}
   m1 := map[string]int{}
   for n1, s1 := range a1 {
      m1[s1] = a2[n1]
   }
   fmt.Println(m1)
}
