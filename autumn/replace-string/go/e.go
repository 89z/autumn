package main
import "strings"

func main() {
   s := "west? east.txt"
   s = strings.Map(func(r rune) rune {
      if strings.ContainsRune(`"*/:<>?\|`, r) { return -1 }
      return r
   }, s)
   println(s == "west east.txt")
}
