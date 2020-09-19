package main
import "fmt"

func main() {
   a := []string{"May", "June", "July"}
   // example 1
   s1 := a[0]
   // example 2
   a2 := a[:2]
   // example 3
   a3 := a[len(a) - 2:]
   // example 4
   s4 := a[len(a) - 1]
   // print
   fmt.Println(s1, a2, a3, s4)
}
