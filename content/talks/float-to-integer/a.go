package main
import "math"
func main() {
   a1 := []string{"Sun", "Mon", "Tue"}
   n1 := math.Round(1.9)
   s1 := a1[n1] // non-integer slice index n1
   println(s1)
}
