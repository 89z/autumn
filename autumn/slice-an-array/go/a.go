package main
import "fmt"

func main() {
   a := []string{"M", "a", "r", "c", "h"}
   // example 1
   s1 := a[2]
   // example 2
   a2 := a[2:4]
   // example 3
   a3 := a[2:]
   // print
   fmt.Println(s1, a2, a3)
}
