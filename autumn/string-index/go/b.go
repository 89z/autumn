package main
import "strings"

func main() {
   n := strings.IndexAny("north", "Oo")
   println(n == 1)
}
