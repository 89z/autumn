package main
import "fmt"

func main() {
   s := "March"
   s = fmt.Sprint(s, "April")
   fmt.Println(s == "MarchApril")
}
