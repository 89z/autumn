package main
import "fmt"

func main() {
   // example 1
   m1 := make(map[string]int)
   // example 2
   m2 := map[string]int{"month": 12, "day": 31}
   // example 3
   m3 := map[string]int{}
   // print
   fmt.Println(m1, m2, m3)
}
