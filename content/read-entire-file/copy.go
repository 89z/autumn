package main
import "io"
import "os"
import "strings"
func main() {
   // head
   var o1, _ = os.Open("a.txt")
   var o2 strings.Builder
   // body
   io.Copy(&o2, o1)
   // foot
   var s1 = o2.String()
   print(s1)
}
