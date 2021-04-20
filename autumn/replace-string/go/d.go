package main
import "strings"

func main() {
   var b strings.Builder
   for _, r := range "west? east.txt" {
      if ! strings.ContainsRune(`"*/:<>?\|`, r) {
         b.WriteRune(r)
      }
   }
   println(b.String() == "west east.txt")
}
