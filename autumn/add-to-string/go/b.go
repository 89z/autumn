package main
import "fmt"

func main() {
   s := "West"
   s = fmt.Sprint(s, "East")
   fmt.Println(s == "WestEast")
}
