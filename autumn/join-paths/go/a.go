package main
import "path/filepath"

func main() {
   var s string

   // example 1
   s = filepath.Join("south", "north")
   println(s == `south\north`)

   // example 2
   s = filepath.Join("south", "north", "")
   println(s == `south\north`)

}
