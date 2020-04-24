package main
import "strings"
func f1(a1... string) string {
   return strings.Join(a1, ",")
}
func main() {
   s1 := f1("Sun", "Mon")
   println(s1)
}
