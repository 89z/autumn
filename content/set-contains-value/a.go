package main
func main() {
   t1 := map[string]bool{"Sun": true, "Mon": true}
   // example 1
   _, b1 := t1["Mon"]
   // example 2
   b2 := t1["Mon"]
   // print
   println(b1, b2)
}
