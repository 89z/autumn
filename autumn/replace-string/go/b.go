package main
import "strings"

func main() {
   var s []string
   for _, r := range `"*/:<>?\|` {
      s = append(s, string(r), "")
   }
   st := strings.NewReplacer(s...).Replace("west? east.txt")
   println(st == "west east.txt")
}
