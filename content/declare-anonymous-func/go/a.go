package main

func main() {
   // example 1
   f1 := func (s string, c byte) bool {
      return s[0] == c
   }
   // example 2
   f2 := func (s, c string) bool {
      return s[:1] == c
   }
   // print
   b1 := f1("June", 'J')
   b2 := f2("June", "J")
   println(b1, b2)
}
