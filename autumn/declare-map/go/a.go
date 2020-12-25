package main
import "fmt"

func main() {
   // example 1
   var m1 = make(map[string]int)
   m1["day"] = 31
   // example 2
   m2 := make(map[string]int)
   m2["day"] = 31
   // print
   fmt.Println(m1, m2)
}
