package main
func main() {
   c1 := map[string]bool{"Sun": true, "Mon": true}
   // example 1
   _, b1 := c1["Mon"]
   // example 2
   b2 := c1["Mon"]
   // print
   println(b1, b2)
}
