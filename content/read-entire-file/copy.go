package main
import "io"
import "os"
import "strings"
func main() {
   var o1, _ = os.Open("a.txt")
   var o2 strings.Builder
   // begin
   io.Copy(&o2, o1)
   // end
   var s1 = o2.String()
   print(s1)
}
