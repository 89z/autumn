package main
import "fmt"

func main() {
   s := "south"
   { // example 1
      t := fmt.Sprintf("%v north", s)
      fmt.Println(t == "south north")
   }
   { // example 2
      t := fmt.Sprintf("%v %[1]v", s)
      fmt.Println(t == "south south")
   }
}
