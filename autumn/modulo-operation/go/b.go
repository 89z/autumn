package main
import "math"

func main() {
   n, n2 := math.Modf(7.5)
   println(n == 7, n2 == .5)
}
