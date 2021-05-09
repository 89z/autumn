package main
import "path"

func main() {
   var s string
   // example 1
   s = path.Join("south", "north")
   println(s == "south/north")
   // example 2
   s = path.Join("south", "north", "")
   println(s == "south/north")
   // example 3
   s = path.Join("https://github.com", "issues")
   println(s == "https:/github.com/issues") // bad
}
