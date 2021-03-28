package main
import "fmt"

func main() {
   s := "West"
   { // example 1
      t := fmt.Sprintf("%v East", s)
      fmt.Println(t == "West East")
   }
   { // example 2
      t := fmt.Sprintf("%v %[1]v", s)
      fmt.Println(t == "West West")
   }
}
