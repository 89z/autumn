package main
import "fmt"

func main() {
   s := "March"
   s2 := fmt.Sprintf("%s,April", s)
   fmt.Println(s2 == "March,April")
}
