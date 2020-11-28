package main
import "math"
type Int int

func (n Int) Add(n2 Int) Int {
   return n + n2
}

func (n Int) Int() int {
   return int(n)
}

func main() {
   // example 1
   n1 := Int.Add(1, 1).Add(1)
   // example 2
   n := Int(1).Add(1).Add(1).Int()
   n2 := math.Pow10(n)
   // print
   println(n1 == 3, n2 == 1000)
}
