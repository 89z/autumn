package main
import "strings"

func main() {
   s := "west? east.txt"
   var a []string
   for _, r := range `"*/:<>?\|` {
      a = append(a, string(r), "")
   }
   s = strings.NewReplacer(a...).Replace(s)
   println(s == "west east.txt")
}
