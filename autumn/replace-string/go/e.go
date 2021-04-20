package main
import "strings"

func main() {
   s := "west? east.txt"
   s = strings.Map(func(r rune) rune {
      switch {
      case strings.ContainsRune(`"*/:<>?\|`, r): return -1
      default: return r
      }
   }, s)
   println(s == "west east.txt")
}
