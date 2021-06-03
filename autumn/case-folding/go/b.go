package main
import "strings"

func main() {
   s := "north"
   s = strings.ToUpper(s)
   println(s == "NORTH")
}
