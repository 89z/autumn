package main
import "fmt"

func main() {
   s := "May"
   s2 := fmt.Sprint(s, "June")
   println(s2 == "MayJune")
}
