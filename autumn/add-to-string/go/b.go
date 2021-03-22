package main
import "fmt"

func main() {
   s := "March"
   { // example 1
      t := fmt.Sprintf("%v April", s)
      fmt.Println(t == "March April")
   }
   { // example 2
      t := fmt.Sprintf("%v %[1]v", s)
      fmt.Println(t == "March March")
   }
}
