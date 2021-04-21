package main
import "strings"

func main() {
   n := strings.IndexByte("north", 'o')
   println(n == 1)
}
