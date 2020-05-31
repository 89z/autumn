package main
import "strings"
func main() {
   s := "♠♣♥♦"
   a := strings.Split(s, "")
   s2 := a[1]
   println(s2 == "♣")
}
