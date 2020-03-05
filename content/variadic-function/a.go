package main
import "strings"
func u1(a1... string) string {
   return strings.Join(a1, " & ")
}
func main() {
   s1 := u1("sun", "mon")
   println(s1)
}
