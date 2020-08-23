package main
import "strings"
func main() {
   s1 := "📕📙📒📗"
   a1 := strings.Split(s1, "")
   s2 := a1[1]
   println(s2 == "📙")
}
