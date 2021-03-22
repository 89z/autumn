package main
import "fmt"

func main() {
   a := []string{"M", "a", "r", "c", "h"}
   { // example 1
      s := a[2]
      fmt.Println(s)
   }
   { // example 2
      b := a[2:4]
      fmt.Println(b)
   }
   { // example 3
      b := a[2:]
      fmt.Println(b)
   }
}
