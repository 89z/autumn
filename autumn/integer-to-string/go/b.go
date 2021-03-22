package main
import "fmt"

func main() {
   n := 100
   { // example 1
      s := fmt.Sprintf("%d", n)
      fmt.Println(s == "100")
   }
   { // example 2
      s := fmt.Sprintf("%v", n)
      fmt.Println(s == "100")
   }
}
