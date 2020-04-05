package main
import "math"
func main() {
   a1 := []string{"Sun", "Mon", "Tue"}
   n1 := 19
   n2 := 10
   n3 := float64(n1) / float64(n2)
   // example 1
   n4 := int(math.Round(n3))
   s1 := a1[n4]
   // example 2
   n5 := int(n3 + .5)
   s2 := a1[n5]
   // example 3
   n6 := int(n3 + math.Copysign(.5, n3))
   s3 := a1[n6]
   // print
   println(s1, s2, s3)
}
