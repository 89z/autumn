package main
import "strings"
func main() {
   a1 := []string{"sun", "mon"}
   s1 := strings.Join(a1, " & ")
   println(s1)
}
