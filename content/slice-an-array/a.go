package main
import "fmt"

func main() {
   a := []string{"May", "June", "July"}
   // example A
   sA := a[0]
   // example B
   aB := a[:2]
   // example C
   aC := a[len(a) - 2:]
   // example D
   sD := a[len(a) - 1]
   // print
   fmt.Println(sA, aB, aC, sD)
}
