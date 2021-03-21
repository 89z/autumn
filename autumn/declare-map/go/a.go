package main
import "fmt"

func main() {
   { // example 1
      var m = make(map[string]int)
      m["day"] = 31
      fmt.Println(m)
   }
   { // example 2
      m := make(map[string]int)
      m["day"] = 31
      fmt.Println(m)
   }
}
