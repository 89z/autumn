package main
import "fmt"
func main() {
   a1 := []string{"Sun", "Mon"}
   for n1 := range a1 {
      a1[n1] += "day"
   }
   fmt.Println(a1)
}
