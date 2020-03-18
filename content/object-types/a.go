package main
import "fmt"
func main() {
   // example 1
   type Jan struct {
      Sun, Mon int
   }
   o1 := Jan{10, 11}
   // example 2
   type Feb struct {
      Sun, Mon, Tue int
   }
   o2 := Feb{Mon: 11, Tue: 12}
   // example 3
   o3 := struct{Sun, Mon int}{10, 11}
   // print
   fmt.Println(o1, o2, o3)
}
