package main
import "strings"
func f1(a1... string) string {
   return strings.Join(a1, ",")
}
var s1 = f1("Sun", "Mon")
func main() {
   println(s1)
}
