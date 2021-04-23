package main
import "strings"

func main() {
   { // example 1
      s := []string{"west", "east"}
      st := strings.Join(s, ",")
      println(st == "west,east")
   }
   { // example 2
      s := []string{"north"}
      st := strings.Join(s, ",")
      println(st == "north")
   }
}
