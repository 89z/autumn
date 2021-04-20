package main
import "strings"

func main() {
   s := "west? east.txt"
   for _, r := range `"*/:<>?\|` {
      s = strings.ReplaceAll(s, string(r), "")
   }
   println(s == "west east.txt")
}
