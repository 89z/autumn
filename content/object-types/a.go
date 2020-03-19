package main
import "fmt"
func main() {
   // normal
   type Jan struct {
      Sun int
      Mon bool
   }
   o1 := Jan{10, true}
   // omit value
   type Feb struct {
      Sun int
      Mon bool
   }
   o2 := Feb{Mon: true}
   // omit key
   type Mar struct {
      int
      bool
   }
   o3 := Mar{10, true}
   // omit type
   o4 := struct {
      Sun int
      Mon bool
   }{10, true}
   // print
   fmt.Println(o1, o2, o3, o4)
}
