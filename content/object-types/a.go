package main
import "fmt"
func main() {
   // example 1
   type jan struct {
      sun, mon int
   }
   o1 := jan{10, 11}
   // example 2
   type feb struct {
      sun, mon, tue int
   }
   o2 := feb{mon: 11, tue: 12}
   // example 3
   o3 := struct{sun, mon int}{10, 11}
   // print
   fmt.Println(o1, o2, o3)
}
