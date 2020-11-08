package main
import "fmt"

func main() {
   s := "March"
   s2 := fmt.Sprint(s, "April")
   fmt.Println(s2 == "MarchApril")
}
