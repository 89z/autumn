package main
import "strings"
func main() {
   a := strings.Split("📕📙📒📗", "")
   s := a[1]
   println(s == "📙")
}
