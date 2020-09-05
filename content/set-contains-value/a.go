package main
func main() {
   t := map[string]bool{"Sunday": true}
   // example 1
   b1, b2 := t["Sunday"]
   // example 2
   b3 := t["Sunday"]
   // print
   println(b1, b2, b3)
}
