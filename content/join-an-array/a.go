package main
import "strings"
func main() {
   a := []string{"Sun", "Mon"}
   s := strings.Join(a, ",")
   println(s == "Sun,Mon")
}
