package main
import "strings"

func main() {
   s := "📗📒📕"
   a := strings.Split(s, "")
   s1 := a[1]
   println(s1 == "📒")
}
