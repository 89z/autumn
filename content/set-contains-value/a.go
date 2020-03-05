package main
func main() {
   c1 := map[string]bool{"sun": true, "mon": true}
   // example 1
   _, b1 := c1["mon"]
   // example 2
   b2 := c1["mon"]
   // print
   println(b1, b2)
}
