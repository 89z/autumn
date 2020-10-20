package main
import "fmt"

func main() {
   s := "May"
   s2 := fmt.Sprintf("%s,June", s)
   println(s2 == "May,June")
}
