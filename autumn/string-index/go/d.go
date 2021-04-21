package main
import "strings"

func main() {
   n := strings.LastIndexByte("west east", 'e')
   println(n == 5)
}
