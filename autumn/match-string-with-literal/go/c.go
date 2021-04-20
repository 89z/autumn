package main
import "strings"

func main() {
   b := strings.ContainsRune(`"*/:<>?\|`, '|')
   println(b)
}
