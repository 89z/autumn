package main
import "strings"

func main() {
   { // example 1
      a := []string{"west", "east"}
      s := strings.Join(a, ",")
      println(s == "west,east")
   }
   { // example 2
      a := []string{"north"}
      s := strings.Join(a, ",")
      println(s == "north")
   }
}
