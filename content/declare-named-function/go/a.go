package main
// example 1
func f1(s string, c byte) bool {
   return s[0] == c
}
// example 2
func f2(s, c string) bool {
   return s[:1] == c
}
// print
func main() {
   b1 := f1("June", 'J')
   b2 := f2("June", "J")
   println(b1, b2)
}
